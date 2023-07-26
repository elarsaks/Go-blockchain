package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/elarsaks/Go-blockchain/pkg/block"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
	"github.com/elarsaks/Go-blockchain/pkg/wallet"
	"github.com/elarsaks/Go-blockchain/wallet_server/handlers"
	"github.com/gorilla/mux"
)

type WalletServer struct {
	port    uint16
	gateway string
}

// Create a new instance of WalletServer
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

// Get the port of the WalletServer
func (ws *WalletServer) Port() uint16 {
	return ws.port
}

// Get the gateway of the WalletServer
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

// Create a new transaction
func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
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
		privateKey := utils.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(*t.Value, 32)
		if err != nil {
			log.Println("ERROR: parse error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		value32 := float32(value)

		w.Header().Add("Content-Type", "application/json")

		transaction := wallet.NewTransaction(privateKey, publicKey,
			*t.SenderBlockchainAddress, *t.RecipientBlockchainAddress, value32)
		signature := transaction.GenerateSignature()
		signatureStr := signature.String()

		bt := &block.TransactionRequest{
			SenderBlockchainAddress:    t.SenderBlockchainAddress,
			RecipientBlockchainAddress: t.RecipientBlockchainAddress,
			SenderPublicKey:            t.SenderPublicKey,
			Value:                      &value32, Signature: &signatureStr,
		}
		m, _ := json.Marshal(bt)
		buf := bytes.NewBuffer(m)

		resp, _ := http.Post(ws.Gateway()+"/transactions", "application/json", buf)

		// Print response
		fmt.Println("response Status:", resp.Status)

		if resp.StatusCode == 201 {
			io.WriteString(w, string(utils.JsonStatus("success")))
			return
		}

		io.WriteString(w, string(utils.JsonStatus("fail")))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

//* Get wallet balance
func (ws *WalletServer) WalletBalance(w http.ResponseWriter, req *http.Request) {
	// Check if the HTTP method is GET
	if req.Method != http.MethodGet {
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract the blockchain address from the URL query parameters
	blockchainAddress := req.URL.Query().Get("blockchain_address")

	// Construct the endpoint URL for the blockchain API
	endpoint := fmt.Sprintf("%s/balance?blockchain_address=%s", ws.Gateway(), blockchainAddress)

	// Send a GET request to the blockchain API
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Printf("ERROR: %v", err)
		io.WriteString(w, string(utils.JsonStatus("fails")))
		return
	}
	defer resp.Body.Close()

	// Set the response header to indicate JSON content type
	w.Header().Set("Content-Type", "application/json")

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		// Decode the response JSON into the existing response struct
		br := &block.BalanceResponse{}
		err := json.NewDecoder(resp.Body).Decode(br)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		// Marshal the response struct to JSON and write it as the response
		m, _ := json.Marshal(br)
		io.WriteString(w, string(m))
	} else {
		// Create a new response struct for the failure case
		failureResponse := &block.BalanceResponse{
			Error: "Failed to get wallet balance",
		}
		m, _ := json.Marshal(failureResponse)
		io.WriteString(w, string(m))
	}
}

// Handler function to get requested blocks
func (ws *WalletServer) GetBlocks(w http.ResponseWriter, req *http.Request) {
	// Get the 'amount' query parameter from the URL
	amountStr := req.URL.Query().Get("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount parameter", http.StatusBadRequest)
		return
	}

	// Make a GET request to miner-2's API to fetch blocks
	resp, err := http.Get(fmt.Sprintf(ws.Gateway()+"/miner/blocks?amount=%d", amount))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error fetching blocks from miner-2", resp.StatusCode)
		return
	}

	// Decode the JSON response into a slice of Block
	var blocks []block.Block
	if err := json.NewDecoder(resp.Body).Decode(&blocks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with JSON-encoded blocks
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Example CORS header, customize as needed
	json.NewEncoder(w).Encode(blocks)
}

// Handler function to get requested blocks
func (ws *WalletServer) GetMinerWallet(w http.ResponseWriter, req *http.Request) {
	// Get the 'miner' query parameter from the URL
	minerID := req.URL.Query().Get("miner_id")

	// TODO: this could be recived from the blockchain (nodes should know each other)
	minerUrl := map[string]string{
		"1": "go-blockchain-miner-1_1:5001",
		"2": "miner-2:5002",
		"3": "miner-3:5003",
	}

	// Make a POST request to the miner's API to fetch the wallet
	requestBody := []byte("optional_request_data")

	fmt.Println("http://" + minerUrl[minerID] + "/miner/wallet")

	resp, err := http.Post(fmt.Sprintf("http://"+minerUrl[minerID]+"/miner/wallet"),
		"application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// TODO: log response
	fmt.Println(resp)

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: Error fetching wallet from %s", minerID)
		http.Error(w, fmt.Sprintf("Error fetching wallet from %s", minerID), resp.StatusCode)
		return
	}

	// Decode the JSON response into a struct or a map
	var walletData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&walletData)
	if err != nil {
		fmt.Printf("Error decoding wallet response")
		http.Error(w, "Error decoding wallet response", http.StatusInternalServerError)
		return
	}

	// Encode the wallet data to JSON and write it to the response
	jsonData, err := json.Marshal(walletData)

	// TODO: log JSON data
	fmt.Println(string(jsonData))
	if err != nil {
		fmt.Printf("Error encoding wallet data")
		http.Error(w, "Error encoding wallet data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Run the WalletServer
func (ws *WalletServer) Run() {
	// Create router
	router := mux.NewRouter()
	router.Use(utils.CorsMiddleware())

	// Map to store route descriptions
	routeDescriptions := map[string]string{
		"/":               "index",
		"/wallet":         "Wallet description...",
		"/wallet/balance": "Wallet balance description...",
		"/transaction":    "Transaction description...",
		"/miner/blocks":   "Miner blocks description...",
	}

	// Return API route descriptions
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routeDescriptions)
	})

	// Example route registration with the WalletServer object
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleExample(w, r, ws.gateway)
	})

	// Register routes
	router.HandleFunc("/wallet", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsetWallet(w, r, ws.gateway)
	})

	router.HandleFunc("/wallet/balance", ws.WalletBalance)
	router.HandleFunc("/transaction", ws.CreateTransaction)
	router.HandleFunc("/miner/blocks", ws.GetBlocks)
	router.HandleFunc("/miner/wallet", ws.GetMinerWallet)

	// Start server
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), router))
}

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {
	// Retrieve gateway from environment variable
	gateway := os.Getenv("WALLET_SERVER_GATEWAY_TO_BLOCKCHAIN")

	if gateway == "" {
		gateway = "http://miner-2:5001" // Default value
	}

	// Retrieve port from environment variable
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 8080 // Default value
	}

	// Create and run the WalletServer with the configured ports and gateway
	app := NewWalletServer(uint16(port), gateway)
	app.Run()
}
