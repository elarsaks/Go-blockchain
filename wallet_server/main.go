package main

import (
	"log"
	"os"
	"strconv"
)

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {
	// Retrieve gateway from environment variable
	gateway := os.Getenv("WALLET_SERVER_GATEWAY")
	if gateway == "" {
		gateway = "http://127.0.0.1:5001" // Default value
	}

	// Retrieve port from environment variable
	portStr := os.Getenv("WALLET_SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 8080 // Default value
	}

	// Print gateway and port
	log.Printf("Gateway: %s\n", gateway)
	log.Printf("Port: %d\n", port)

	app := NewWalletServer(uint16(port), gateway)
	app.Run()
}
