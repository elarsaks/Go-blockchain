package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/elarsaks/Go-blockchain/pkg/block"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
	"github.com/elarsaks/Go-blockchain/pkg/wallet"
)

// Handler function to get requested blocks
func GetMinerWallet(w http.ResponseWriter, req *http.Request, miner string) {
	// Get the 'miner' query parameter from the URL
	minerID := req.URL.Query().Get("miner_id")

	// TODO: this could be recived from the blockchain (nodes should know each other)
	minerUrl := map[string]string{
		"1": "miner-1:5001",
		"2": "miner-2:5002",
		"3": "miner-3:5003",
	}

	// Make a POST request to the miner's API to fetch the wallet
	requestBody := []byte("optional_request_data")

	resp, err := http.Post(fmt.Sprintf("http://"+minerUrl[minerID]+"/miner/wallet"),
		"application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

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

	if err != nil {
		fmt.Printf("Error encoding wallet data")
		http.Error(w, "Error encoding wallet data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Get User wallet
// TODO: Refactor this function (It creates a wallet and registers it on the blockchain)
func GetUserWallet(w http.ResponseWriter, req *http.Request, miner string) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid HTTP Method", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	userWallet := wallet.NewWallet()

	// Create a payload containing the userWallet's blockchain address
	payload := struct {
		BlockchainAddress string `json:"blockchainAddress"`
	}{
		BlockchainAddress: userWallet.BlockchainAddress(),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("ERROR: Failed to marshal payload:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Register the userWallet on the blockchain
	resp, err := http.Post(miner+"/wallet/register", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("ERROR: Failed to register wallet: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("ERROR: Failed to register wallet")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the userWallet as part of the response
	userWalletBytes, err := json.Marshal(userWallet)
	if err != nil {
		log.Println("ERROR: Failed to marshal userWallet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: return error messages
	io.WriteString(w, string(userWalletBytes))
}

//* Get wallet balance
func GetWalletBalance(w http.ResponseWriter, req *http.Request, miner string) {
	// Check if the HTTP method is GET
	if req.Method != http.MethodGet {
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract the blockchain address from the URL query parameters
	blockchainAddress := req.URL.Query().Get("blockchain_address")

	// Construct the endpoint URL for the blockchain API
	endpoint := fmt.Sprintf("%s/balance?blockchain_address=%s", miner, blockchainAddress)

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
