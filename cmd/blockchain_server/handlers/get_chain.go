package handlers

// // Get the gateway of the BlockchainServer
// func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodGet:
// 		w.Header().Add("Content-Type", "application/json")
// 		bc := bcs.GetBlockchain()
// 		m, _ := bc.MarshalJSON()
// 		io.WriteString(w, string(m[:]))
// 	default:
// 		log.Printf("ERROR: Invalid HTTP Method")

// 	}
// }
