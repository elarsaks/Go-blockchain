package handlers

// Get the wallet of the BlockchainServer // NOTE: This is not a part of the blockchain
// func (bcs *BlockchainServer) MinerWallet(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodPost:
// 		w.Header().Add("Content-Type", "application/json")
// 		myWallet := bcs.Wallet
// 		m, _ := myWallet.MarshalJSON()
// 		io.WriteString(w, string(m[:]))
// 	default:
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println("ERROR: Invalid HTTP Method")
// 	}
// }
