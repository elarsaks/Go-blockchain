package handlers

import (
	"io"
	"log"
	"net/http"
)

// TODO: Describe the purpose of this function

// Get the gateway of the BlockchainServer
func (h *BlockchainServerHandler) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := h.server.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")

	}
}
