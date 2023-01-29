package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Define Block type
type Block struct {
	nonce        int
	previousHash [32]byte
	transactions []string
	timeStamp    int64
}

// Create new block //* Note: this is not called anywhere
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timeStamp = time.Now().UnixNano()
	return b
}

// Print the block
func (b *Block) Print() {
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %x\n", b.previousHash)
	fmt.Printf("transactions    %s\n", b.transactions)
	fmt.Printf("time_stamp      %d\n", b.timeStamp)
}

// Define Blockchain type
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// Create blockchain (including genesis block)
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// Create new block
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := &Block{}
	bc.chain = append(bc.chain, b)
	return b
}

// Get last block
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// Print blockchain
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

// Generate sha256 hash from a block, for a block
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)

	fmt.Println(string(m))

	return sha256.Sum256([]byte(m))
}

// Marshal block to JSON
//TODO: Learn more about this
// This function should help return the JSON representation of the block
// It should translate struct to JSON
func (b *Block) MashalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
		TimeStamp    int64    `json:"time_stamp"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
		TimeStamp:    b.timeStamp,
	})
}

// Function to initialize the logger
func init() { //? TODO: Why we need this?
	log.SetPrefix("Blockchain: ")
}

// Main function
func main() {
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())

	// Initialize a new blockchain ()
	// blockChain := NewBlockchain()
	// blockChain.Print()

	// Create a new block
	// previousHash := blockChain.LastBlock().Hash()
	// blockChain.CreateBlock(5, previousHash)
	// blockChain.Print()

	// Create a new block
	// previousHash = blockChain.LastBlock().Hash()
	// blockChain.CreateBlock(2, previousHash)
	// blockChain.Print()
}
