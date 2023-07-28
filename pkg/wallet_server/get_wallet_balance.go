package wallet_server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/elarsaks/Go-blockchain/pkg/block"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

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
