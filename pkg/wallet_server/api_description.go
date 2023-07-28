package wallet_server

import (
	"encoding/json"
	"net/http"
)

func (ws *WalletServer) GetApiDescription(w http.ResponseWriter, r *http.Request) {
	description := map[string]string{
		"/":               "index",
		"/wallet":         "Wallet description...",
		"/wallet/balance": "Wallet balance description...",
		"/transaction":    "Transaction description...",
		"/miner/blocks":   "Miner blocks description...",
	}

	// convert map to json
	jsonData, err := json.Marshal(description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// write the json to the response body
	w.Write(jsonData)
}
