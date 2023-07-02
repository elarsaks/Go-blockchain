package main

import (
	"log"
	"os"
	"strconv"
)

// Function to initialize the logger
func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// Retrieve port from environment variable
	portStr := os.Getenv("BLOCKCHAIN_SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 5001 // Default value
	}

	// Print port
	log.Printf("Port: %d\n", port)

	app := NewBlockchainServer(uint16(port))
	app.Run()
}
