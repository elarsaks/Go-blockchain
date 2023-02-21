package main

import (
	"Go-blockchain/block"
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
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet A sends 1.0 to Wallet B
	t := wallet.NewTransaction(
		walletA.PrivateKey(),
		walletA.PublicKey(),
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0)

	// Blockchain Node (Miner)
	blockchain := block.NewBlockchain(walletM.BlockchainAddress())

	// Wallet A sends 1.0 to Wallet B (Transaction)
	isAdded := blockchain.AddTransaction(
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0,
		walletA.PublicKey(),
		t.GeneratreSignature())

	fmt.Println("Added: ", isAdded)
}
