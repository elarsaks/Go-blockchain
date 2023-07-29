package handlers

// import (
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/elarsaks/Go-blockchain/pkg/utils"
// )

// // Resolve the conflicts in the BlockchainServer
// func (h *BlockchainServerHandler) Consensus(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodPut:
// 		bc := h.server.GetBlockchain()
// 		replaced := bc.ResolveConflicts()

// 		w.Header().Add("Content-Type", "application/json")
// 		if replaced {
// 			io.WriteString(w, string(utils.JsonStatus("success")))
// 		} else {
// 			io.WriteString(w, string(utils.JsonStatus("fail")))
// 		}
// 	default:
// 		log.Printf("ERROR: Invalid HTTP Method")
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }
