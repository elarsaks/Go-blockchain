package main

import (
	"Go-blockchain/wallet"
	"fmt"
	"log"
)

// Function to initialize the logger
func init() {
	log.SetPrefix("Blockchain: ")
}

// Main function
func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())
}
