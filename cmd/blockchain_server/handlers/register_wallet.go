package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Register the wallet in the BlockchainServer
func (h *BlockchainServerHandler) RegisterWallet(w http.ResponseWriter, req *http.Request) {
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
		success := h.server.GetBlockchain().RegisterNewWallet(requestBody.BlockchainAddress, "Register User Wallet")
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
