package main

import (
	"fmt"
	"log"
)

// Function to initialize the logger
func init() {
	log.SetPrefix("Blockchain: ")
}

// Main function
func main() {
	w := NewWallet()
	fmt.Println(w.privateKeyStr())
	fmt.Println(w.publicKeyStr())
}
