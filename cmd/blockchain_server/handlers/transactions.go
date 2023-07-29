package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/elarsaks/Go-blockchain/pkg/block"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

// TODO: Divide this function into 4 different functions

// Transactions is a handler function that allows for getting, creating, updating and deleting transactions
func (h *BlockchainServerHandler) Transactions(w http.ResponseWriter, req *http.Request) {
	// Switching on the HTTP method
	switch req.Method {

	// In case of a GET request, return the current transaction pool
	case http.MethodGet:
		// Setting the Content-Type of the response to application/json
		w.Header().Add("Content-Type", "application/json")

		// Getting the blockchain from the server
		bc := h.server.GetBlockchain()

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
		bc := h.server.GetBlockchain()

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
		bc := h.server.GetBlockchain()

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
		bc := h.server.GetBlockchain()

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
