package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/elarsaks/Go-blockchain/pkg/block"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
	"github.com/elarsaks/Go-blockchain/pkg/wallet"
)

// CreateTransaction is responsible for creating a new transaction
func CreateTransaction(w http.ResponseWriter, req *http.Request, miner string) {
	// Switching on the HTTP method
	switch req.Method {
	case http.MethodPost:

		// Decoding the body of the request into a TransactionRequest object
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
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

		// Convert the sender's public and private keys from strings to their appropriate types
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)

		// Parse the value from the request, handle error if the value is not a valid float
		value, err := strconv.ParseFloat(*t.Value, 32)
		if err != nil {
			log.Println("ERROR: parse error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		value32 := float32(value)

		// Setting the Content-Type of the response to application/json
		w.Header().Add("Content-Type", "application/json")

		// Create a new Transaction object
		transaction := wallet.NewTransaction(privateKey,
			publicKey,
			*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress,
			*t.Message,
			value32)

		// Generate a signature for the transaction
		signature := transaction.GenerateSignature()
		signatureStr := signature.String()

		// Create a new TransactionRequest object that will be sent to the miner
		bt := &block.TransactionRequest{
			SenderBlockchainAddress:    t.SenderBlockchainAddress,
			RecipientBlockchainAddress: t.RecipientBlockchainAddress,
			SenderPublicKey:            t.SenderPublicKey,
			Message:                    t.Message,
			Value:                      &value32,
			Signature:                  &signatureStr,
		}

		// Serialize the TransactionRequest object into JSON
		m, _ := json.Marshal(bt)
		buf := bytes.NewBuffer(m)

		// Make a POST request to the miner's API to create a new transaction
		resp, err := http.Post(miner+"/transactions", "application/json", buf)

		// Check if there was an error while making the POST request
		if err != nil {
			// Log the error message
			log.Printf("ERROR: Failed to make POST request: %v", err)

			// Pass the error message to the client
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// Log the error message
			log.Printf("ERROR: Failed to read response body: %v", err)

			// Pass the error message to the client
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check the response status code and send a success response if it was 201
		if resp.StatusCode == 201 {
			io.WriteString(w, string(utils.JsonStatus("success")))
			return
		}

		// If the status code was not 201, send the response body (which contains the error message) to the client
		w.WriteHeader(resp.StatusCode)
		io.WriteString(w, string(body))

	// If the HTTP method is not POST, send a 400 response and log an error message
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

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
