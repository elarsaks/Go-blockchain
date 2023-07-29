package handlers

// // Start the mining process in the BlockchainServer
// func (bcs *BlockchainServer) StartMine(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodGet:
// 		bc := bcs.GetBlockchain()
// 		bc.StartMining()

// 		m := utils.JsonStatus("success")
// 		w.Header().Add("Content-Type", "application/json")
// 		io.WriteString(w, string(m))
// 	default:
// 		log.Println("ERROR: Invalid HTTP Method")
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }
