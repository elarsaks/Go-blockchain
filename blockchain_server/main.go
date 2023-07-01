package main

import (
	"flag"
	"log"
)

// Function to initialize the logger
func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// TODO: Port and should come from .env file
	port := flag.Uint("port", 5001, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
