package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetMinerWallet is a handler that:
// 1. Gets the 'miner' query parameter from the URL
// 2. Sets the wallet server's gateway to blockhain to the requested miner
// 3. Makes a POST request to the miner's API to fetch the wallet
// 4. Checks the response status code
// 5. Decodes the JSON response into a struct or a map
// 6. Encodes the wallet data to JSON and writes it to the response
func (h *WalletServerHandler) GetMinerWallet(w http.ResponseWriter, req *http.Request) {
	// Get the 'miner' query parameter from the URL
	minerID := req.URL.Query().Get("miner_id")

	// Set the gateway to the miner
	h.server.SetGateway(minerID)
	// Make a POST request to the miner's API to fetch the wallet
	requestBody := []byte("optional_request_data")

	resp, err := http.Post(fmt.Sprintf(h.server.Gateway()+"/miner/wallet"),
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
