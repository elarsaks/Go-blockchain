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

	"github.com/elarsaks/Go-blockchain/block"
	"github.com/elarsaks/Go-blockchain/utils"
	"github.com/elarsaks/Go-blockchain/wallet"
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

// Get the wallet of the WalletServer
func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
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

// Run the WalletServer
func (ws *WalletServer) Run() {
	// Create router
	router := mux.NewRouter()
	router.Use(utils.CorsMiddleware())

	// Define routes
	router.HandleFunc("/wallet", ws.Wallet)
	router.HandleFunc("/wallet/balance", ws.WalletBalance)
	router.HandleFunc("/transaction", ws.CreateTransaction)

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
		gateway = "http://127.0.0.1:5001" // Default value
	}

	// Retrieve port from environment variable
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 5000 // Default value
	}

	// Print gateway and port
	log.Printf("Gateway to blockchain: %s\n", gateway)
	log.Printf("Port: %d\n", port)

	app := NewWalletServer(uint16(port), gateway)
	app.Run()
}
