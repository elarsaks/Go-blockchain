package handlers

// import (
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/elarsaks/Go-blockchain/pkg/utils"
// )

// // Mine the Block in the BlockchainServer
// func (bcs *BlockchainServer) Mine(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodGet:
// 		bc := bcs.GetBlockchain()
// 		isMined := bc.Mining()

// 		var m []byte
// 		if !isMined {
// 			w.WriteHeader(http.StatusBadRequest)
// 			m = utils.JsonStatus("fail")
// 		} else {
// 			m = utils.JsonStatus("success")
// 		}
// 		w.Header().Add("Content-Type", "application/json")
// 		io.WriteString(w, string(m))
// 	default:
// 		log.Println("ERROR: Invalid HTTP Method")
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }
