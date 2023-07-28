package wallet_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/elarsaks/Go-blockchain/pkg/wallet"
)

// Get User wallet
func (ws *WalletServer) GetUserWallet(w http.ResponseWriter, req *http.Request) {

	fmt.Println("TESTING!")

	if req.Method != http.MethodPost {
		http.Error(w, "Invalid HTTP Method", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	userWallet := wallet.NewWallet()

	// Create a payload containing the userWallet's blockchain address
	payload := struct {
		BlockchainAddress string `json:"blockchainAddress"`
	}{
		BlockchainAddress: userWallet.BlockchainAddress(),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("ERROR: Failed to marshal payload:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Register the userWallet on the blockchain
	resp, err := http.Post(ws.Gateway()+"/wallet/register", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("ERROR: Failed to register wallet: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("ERROR: Failed to register wallet")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the userWallet as part of the response
	userWalletBytes, err := json.Marshal(userWallet)
	if err != nil {
		log.Println("ERROR: Failed to marshal userWallet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: return error messages
	io.WriteString(w, string(userWalletBytes))
}
