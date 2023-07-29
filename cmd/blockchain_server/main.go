package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

// Get the wallet balance by blockchain address in the Blockchain
func (bcs *BlockchainServer) Balance(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		br := &block.BalanceResponse{} // Use the BalanceResponse type

		blockchainAddress := req.URL.Query().Get("blockchainAddress")

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

// // Get the last 10 blocks of the BlockchainServer
// func (bcs *BlockchainServer) GetBlocks(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodGet:
// 		w.Header().Add("Content-Type", "application/json")
// 		bc := bcs.GetBlockchain()
// 		m, _ := json.Marshal(bc.GetBlocks(10))
// 		io.WriteString(w, string(m[:]))
// 	default:
// 		log.Printf("ERROR: Invalid HTTP Method")
// 	}
// }

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

// Get the blockchain of the BlockchainServer
func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc

		//? Why do we need to register the wallet here ?
		// Call RegisterNewWallet to register the provided wallet address
		success := bc.RegisterNewWallet(minersWallet.BlockchainAddress(), "Register Miner Wallet")
		if !success {
			log.Println("ERROR: Failed to register wallet")
			// TODO: Handle error
			return nil
		}

		// Setting the wallet in the BlockchainServer object
		bcs.Wallet = minersWallet

		log.Printf("privateKey %v", minersWallet.PrivateKeyStr())
		log.Printf("publicKey %v", minersWallet.PublicKeyStr())
		log.Printf("blockchainAddress %v", minersWallet.BlockchainAddress())
	}
	return bc
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

// Create a new instance of BlockchainServer
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{
		port:   port,
		Wallet: nil,
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
		success := bcs.GetBlockchain().RegisterNewWallet(requestBody.BlockchainAddress, "Register User Wallet")
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

// TODO: Divide this function into 4 different functions
// Transactions is a handler function that allows for getting, creating, updating and deleting transactions
func (bcs *BlockchainServer) Transactions(w http.ResponseWriter, req *http.Request) {
	// Switching on the HTTP method
	switch req.Method {

	// In case of a GET request, return the current transaction pool
	case http.MethodGet:
		// Setting the Content-Type of the response to application/json
		w.Header().Add("Content-Type", "application/json")

		// Getting the blockchain from the server
		bc := bcs.GetBlockchain()

		// Getting the transaction pool
		transactions := bc.TransactionPool()

		// Serializing the transaction pool and its length into a JSON object
		m, _ := json.Marshal(struct {
			Transactions []*block.Transaction `json:"transactions"`
			Length       int                  `json:"length"`
		}{
			Transactions: transactions,
			Length:       len(transactions),
		})

		// Writing the serialized JSON object to the response
		io.WriteString(w, string(m[:]))

	// In case of a POST request, create a new transaction
	case http.MethodPost:
		// Decoding the body of the request into a TransactionRequest object
		decoder := json.NewDecoder(req.Body)
		var t block.TransactionRequest
		err := decoder.Decode(&t)

		// If there was an error decoding the request, log the error and send a fail response
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// Validate the transaction request, send a fail response if validation fails
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// Convert the sender's public key and signature from strings to their appropriate types
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)

		// Getting the blockchain from the server
		bc := bcs.GetBlockchain()

		// Attempting to create a new transaction
		isCreated, err := bc.CreateTransaction(*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress, *t.Message, *t.Value, publicKey, signature)

		// Setting the Content-Type of the response to application/json
		w.Header().Add("Content-Type", "application/json")

		var m []byte
		if err != nil { // If there is an error during the transaction creation

			w.WriteHeader(http.StatusBadRequest)
			// Create an anonymous struct to hold the status and the error message
			errMsg := struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{
				Status:  "fail",
				Message: err.Error(),
			}

			// Marshal the struct into a JSON string
			m, _ = json.Marshal(errMsg)
		} else if !isCreated { // If the transaction was not created successfully
			fmt.Println(" If the transaction was not created successfully")

			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else { // If the transaction was created successfully
			w.WriteHeader(http.StatusCreated)
			m = utils.JsonStatus("success")
		}

		// Write the status message to the response
		io.WriteString(w, string(m))

	// In case of a PUT request, update a transaction
	case http.MethodPut:
		// Decoding the body of the request into a TransactionRequest object
		decoder := json.NewDecoder(req.Body)
		var t block.TransactionRequest
		err := decoder.Decode(&t)

		// If there was an error decoding the request, log the error and send a fail response
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// Validate the transaction request, send a fail response if validation fails
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// Convert the sender's public key and signature from strings to their appropriate types
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)

		// Getting the blockchain from the server
		bc := bcs.GetBlockchain()

		// Updating a transaction and getting whether the transaction was updated successfully

		isUpdated, err := bc.AddTransaction(*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress,
			*t.Message,
			*t.Value,
			publicKey,
			signature)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"status": "fail", "error": err.Error()})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"status": "success"})
		}

		// Setting the Content-Type of the response to application/json
		w.Header().Add("Content-Type", "application/json")

		// If the transaction was not updated successfully, send a fail status message
		var m []byte
		if !isUpdated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail")
		} else { // If the transaction was updated successfully, send a success status message
			m = utils.JsonStatus("success")
		}

		// Write the status message to the response
		io.WriteString(w, string(m))

	// In case of a DELETE request, clear the transaction pool
	case http.MethodDelete:
		// Getting the blockchain from the server
		bc := bcs.GetBlockchain()

		// Clearing the transaction pool
		bc.ClearTransactionPool()

		// Send a success status message
		io.WriteString(w, string(utils.JsonStatus("success")))

	// In case of an unsupported HTTP method, log an error and send a 400 response
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

	handler := handlers.NewBlockchainServerHandler(bcs)

	router.HandleFunc("/miner/blocks", handler.GetBlocks)
	router.HandleFunc("/miner/wallet", handler.MinerWallet)

	// Define routes
	// router.HandleFunc("/miner/wallet", bcs.MinerWallet)
	router.HandleFunc("/wallet/register", bcs.RegisterWallet)
	router.HandleFunc("/", bcs.GetChain)
	// router.HandleFunc("/miner/blocks", bcs.GetBlocks)
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
