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

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "recipient", 1.0)
	fmt.Printf("signature: %s\n", t.GeneratreSignature())
}
