package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Get the last 10 blocks of the BlockchainServer
func (h *BlockchainServerHandler) GetBlocks(w http.ResponseWriter, req *http.Request) {

	fmt.Println("GET _ _ _ _ _ _ _ BLOCKS")

	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		// bc := h.server.GetBlockchain()
		//m, _ := json.Marshal(bc.GetBlocks(10))
		// io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}
