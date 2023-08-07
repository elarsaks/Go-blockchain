package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/elarsaks/Go-blockchain/cmd/blockchain_server/handlers"
	"github.com/elarsaks/Go-blockchain/pkg/block"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
	"github.com/elarsaks/Go-blockchain/pkg/wallet"
	"github.com/gorilla/mux"
)

// TODO: Divide this file into multiple files

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

type BlockchainServer struct {
	port   uint16
	Wallet *wallet.Wallet
	//* NOTE: In real world app we would not attach the wallet to the server
	// But for the sake of simplicity we will do it here,
	// because we dont store miners credentials in a database.
}

// Get the port of the BlockchainServer
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

// GetWallet method for BlockchainServer
func (bcs *BlockchainServer) GetWallet() *wallet.Wallet {
	return bcs.Wallet
}

// Create a new instance of BlockchainServer
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{
		port:   port,
		Wallet: nil,
	}
}

// Get the blockchain of the BlockchainServer
func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc

		//* DEBUG #Consensus
		//? Why do we need to register the wallet here ?
		/* Call RegisterNewWallet to register the provided wallet address
		success := bc.RegisterNewWallet(minersWallet.BlockchainAddress(), "Register Miner Wallet")
		if !success {
			log.Println("ERROR: Failed to register wallet")
			// TODO: Handle error
			return nil
		}

		// Setting the wallet in the BlockchainServer object
		bcs.Wallet = minersWallet */

		// log.Printf("privateKey %v", minersWallet.PrivateKeyStr())
		// log.Printf("publicKey %v", minersWallet.PublicKeyStr())
		// log.Printf("blockchainAddress %v", minersWallet.BlockchainAddress())
	}
	return bc
}

// RegisterMinersWallet registers the miner's wallet
func (bcs *BlockchainServer) RegisterMinersWallet() {
	minersWallet := wallet.NewWallet()

	portStr := strconv.Itoa(int(bcs.Port()))
	lastDigit := portStr[len(portStr)-1:]
	message := "Register Miner Wallet " + lastDigit

	// Call RegisterNewWallet to register the provided wallet address
	success := bcs.GetBlockchain().RegisterNewWallet(minersWallet.BlockchainAddress(), message)
	if !success {
		log.Println("ERROR: Failed to register wallet")
		// TODO: Handle error
		return
	}

	// Setting the wallet in the BlockchainServer object
	bcs.Wallet = minersWallet

	log.Printf("privateKey %v", minersWallet.PrivateKeyStr())
	log.Printf("publicKey %v", minersWallet.PublicKeyStr())
	log.Printf("blockchainAddress %v", minersWallet.BlockchainAddress())
}

// Run the BlockchainServer
func (bcs *BlockchainServer) Run() {
	bcs.GetBlockchain().Run()

	// Register the miner's wallet
	bcs.RegisterMinersWallet()

	router := mux.NewRouter()
	router.Use(utils.CorsMiddleware())

	handler := handlers.NewBlockchainServerHandler(bcs)

	// Define routes
	router.HandleFunc("/chain", handler.GetChain)
	router.HandleFunc("/balance", handler.Balance)
	router.HandleFunc("/consensus", handler.Consensus)
	router.HandleFunc("/mine", handler.Mine)
	router.HandleFunc("/mine/start", handler.StartMine)
	router.HandleFunc("/miner/blocks", handler.GetBlocks)
	router.HandleFunc("/miner/wallet", handler.MinerWallet)
	router.HandleFunc("/transactions", handler.Transactions)
	router.HandleFunc("/wallet/register", handler.RegisterWallet)

	// Start the server
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), router))
}

// Function to initialize the logger
func init() {
	log.SetPrefix("Blockchain: ")
}

// Main function
func main() {
	// Retrieve port from environment variable
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 5001 // Default value
	}

	// Print port
	log.Printf("Port: %d\n", port)

	app := NewBlockchainServer(uint16(port))
	app.Run()
}
