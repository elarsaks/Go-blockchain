package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

// TODO: Desciption

// Start the mining process in the BlockchainServer
func (h *BlockchainServerHandler) StartMine(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := h.server.GetBlockchain()
		bc.StartMining()

		m := utils.JsonStatus("success")
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m))
	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}
