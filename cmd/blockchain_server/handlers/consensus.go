package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

// Consensus handles the consensus mechanism for the blockchain.
// It resolves any conflicts in the BlockchainServer to ensure data integrity.
func (h *BlockchainServerHandler) Consensus(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPut:
		// Get the current blockchain instance from the server
		bc := h.server.GetBlockchain()

		// Attempt to resolve any conflicts in the blockchain
		replaced := bc.ResolveConflicts()

		// Set the response header to indicate JSON content
		w.Header().Add("Content-Type", "application/json")

		// If conflicts were resolved and the chain was replaced, return "success"
		// Otherwise, return "fail"
		if replaced {
			io.WriteString(w, string(utils.JsonStatus("success")))
		} else {
			io.WriteString(w, string(utils.JsonStatus("fail")))
		}
	default:
		// Log an error if an unsupported HTTP method is used
		log.Printf("ERROR: Invalid HTTP Method")

		// Return a 400 Bad Request status to the client
		w.WriteHeader(http.StatusBadRequest)
	}
}
