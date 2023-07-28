package main

import (
	"log"
	"os"
	"strconv"

	"github.com/elarsaks/Go-blockchain/pkg/wallet_server"
)

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {
	// Retrieve gateway from environment variable
	gateway := os.Getenv("WALLET_SERVER_GATEWAY_TO_BLOCKCHAIN")

	if gateway == "" {
		gateway = "http://127.0.0.1:5001" // Default value
	}

	// Retrieve port from environment variable
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 5000 // Default value
	}

	// Print gateway and port
	log.Printf("Gateway to blockchain: %s\n", gateway)
	log.Printf("Port: %d\n", port)

	app := wallet_server.NewWalletServer(uint16(port), gateway)
	app.Run()
}
