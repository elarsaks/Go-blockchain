package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Example handler that accesses the WalletServer
func HandleExample(w http.ResponseWriter, r *http.Request, gateway string) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)

	// Access and use the WalletServer object
	log.Println("Gateway:", gateway)
}
