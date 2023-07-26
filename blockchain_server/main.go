package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/elarsaks/Go-blockchain/pkg/block"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
	"github.com/elarsaks/Go-blockchain/pkg/wallet"
	"github.com/gorilla/mux"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

type BlockchainServer struct {
	port   uint16
	Wallet *wallet.Wallet
	//* NOTE: In real world app we would not attach the wallet to the server
	// But for the sake of simplicity we will do it here,
	// because we dont store miners credentials in a database.
}

// Create a new instance of BlockchainServer
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{
		port:   port,
		Wallet: nil,
	}
}

// Get the port of the BlockchainServer
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

// Get the blockchain of the BlockchainServer
func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc

		// Setting the wallet in the BlockchainServer object
		bcs.Wallet = minersWallet

		// Call RegisterNewWallet to register the provided wallet address
		success := bc.RegisterNewWallet(minersWallet.BlockchainAddress())
		if !success {
			log.Println("ERROR: Failed to register wallet")
			// TODO: Handle error
			return nil
		}

		log.Printf("private_key %v", minersWallet.PrivateKeyStr())
		log.Printf("public_key %v", minersWallet.PublicKeyStr())
		log.Printf("blockchain_address %v", minersWallet.BlockchainAddress())
	}
	return bc
}

// Get the gateway of the BlockchainServer
func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")

	}
}

// Get the last 10 blocks of the BlockchainServer
func (bcs *BlockchainServer) GetBlocks(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := json.Marshal(bc.GetBlocks(10))
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

// Get the transactions of the BlockchainServer
func (bcs *BlockchainServer) Transactions(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		transactions := bc.TransactionPool()
		m, _ := json.Marshal(struct {
			Transactions []*block.Transaction `json:"transactions"`
			Length       int                  `json:"length"`
		}{
			Transactions: transactions,
			Length:       len(transactions),
		})
		io.WriteString(w, string(m[:]))

	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t block.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)
		bc := bcs.GetBlockchain()
		isCreated := bc.CreateTransaction(*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress, *t.Value, publicKey, signature)

		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isCreated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else {
			w.WriteHeader(http.StatusCreated)
			m = utils.JsonStatus("success")
		}
		io.WriteString(w, string(m))

	case http.MethodPut:
		decoder := json.NewDecoder(req.Body)
		var t block.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)
		bc := bcs.GetBlockchain()
		isUpdated := bc.AddTransaction(*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress, *t.Value, publicKey, signature)

		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isUpdated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else {
			m = utils.JsonStatus("success")
		}
		io.WriteString(w, string(m))

	case http.MethodDelete:
		bc := bcs.GetBlockchain()
		bc.ClearTransactionPool()
		io.WriteString(w, string(utils.JsonStatus("success")))

	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Mine the Block in the BlockchainServer
func (bcs *BlockchainServer) Mine(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		isMined := bc.Mining()

		var m []byte
		if !isMined {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else {
			m = utils.JsonStatus("success")
		}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m))
	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Start the mining process in the BlockchainServer
func (bcs *BlockchainServer) StartMine(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		bc.StartMining()

		m := utils.JsonStatus("success")
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m))
	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Get the wallet balance by blockchain address in the Blockchain
func (bcs *BlockchainServer) Balance(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		br := &block.BalanceResponse{} // Use the BalanceResponse type

		blockchainAddress := req.URL.Query().Get("blockchain_address")
		balance, err := bcs.GetBlockchain().CalculateTotalBalance(blockchainAddress)

		br.Balance = balance
		if err != nil {
			log.Printf("ERROR: %v", err)
			br.Error = err.Error()
		}

		m, _ := json.Marshal(br)

		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(m))

	default:
		log.Printf("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Resolve the conflicts in the BlockchainServer
func (bcs *BlockchainServer) Consensus(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPut:
		bc := bcs.GetBlockchain()
		replaced := bc.ResolveConflicts()

		w.Header().Add("Content-Type", "application/json")
		if replaced {
			io.WriteString(w, string(utils.JsonStatus("success")))
		} else {
			io.WriteString(w, string(utils.JsonStatus("fail")))
		}
	default:
		log.Printf("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Get the wallet of the BlockchainServer // NOTE: This is not a part of the blockchain
func (bcs *BlockchainServer) MinerWallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := bcs.Wallet
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

// Register the wallet in the BlockchainServer
func (bcs *BlockchainServer) RegisterWallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		// Read the request body
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Restore the request body to its original state
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Define a struct to capture the request body
		type RequestBody struct {
			BlockchainAddress string `json:"blockchainAddress"`
		}

		// Decode the request body into the struct
		var requestBody RequestBody
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			log.Println("ERROR: Failed to decode request body:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Call RegisterNewWallet to register the provided wallet address
		success := bcs.GetBlockchain().RegisterNewWallet(requestBody.BlockchainAddress)
		if !success {
			log.Println("ERROR: Failed to register wallet")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Optionally, you can return a success response
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Wallet registered successfully",
		}
		m, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(m)

	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Run the BlockchainServer
func (bcs *BlockchainServer) Run() {
	bcs.GetBlockchain().Run()

	router := mux.NewRouter()
	router.Use(utils.CorsMiddleware())

	// Define routes
	router.HandleFunc("/miner/wallet", bcs.MinerWallet)
	router.HandleFunc("/wallet/register", bcs.RegisterWallet)
	router.HandleFunc("/", bcs.GetChain)
	router.HandleFunc("/miner/blocks", bcs.GetBlocks)
	router.HandleFunc("/transactions", bcs.Transactions)
	router.HandleFunc("/mine", bcs.Mine)
	router.HandleFunc("/mine/start", bcs.StartMine)
	router.HandleFunc("/balance", bcs.Balance)
	router.HandleFunc("/consensus", bcs.Consensus)

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
