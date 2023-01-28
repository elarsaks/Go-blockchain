package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

//Define Block type
type Block struct {
	nonce        int
	previousHash string
	transactions []string
	timeStamp    int64
}

// Function to create a new block
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timeStamp = time.Now().UnixNano()
	return b
}

// Function to print the block
func (b *Block) Print() {
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %s\n", b.previousHash)
	fmt.Printf("transactions    %s\n", b.transactions)
	fmt.Printf("time_stamp      %d\n", b.timeStamp)
}

// Define Blockchain type
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// Function to create a new blockchain
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

// Function to create and add a new block to the blockchain
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// Function to print the blockchain
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

// Function to initialize the logger
func init() {
	log.SetPrefix("Blockchain: ")
}

// Main function
func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash2")
	blockChain.Print()
}
