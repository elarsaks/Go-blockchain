package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/elarsaks/Go-blockchain/pkg/block"
)

// Handler function to get requested blocks
func GetBlocks(w http.ResponseWriter, req *http.Request, miner string) {
	// Get the 'amount' query parameter from the URL
	amountStr := req.URL.Query().Get("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount parameter", http.StatusBadRequest)
		return
	}

	// Make a GET request to miner-2's API to fetch blocks
	resp, err := http.Get(fmt.Sprintf(miner+"/miner/blocks?amount=%d", amount))
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
