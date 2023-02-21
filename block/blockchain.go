package block

import (
	"Go-blockchain/utils"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const MINING_DIFFICULTY = 3
const MINING_SENDER = "THE BLOCKCHAIN"
const MINING_REWARD = 1.0

// Define Block type
type Block struct {
	timeStamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

// Create new block
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timeStamp = time.Now().UnixNano()
	b.transactions = transactions
	return b
}

// Print the block
func (b *Block) Print() {
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %x\n", b.previousHash)
	fmt.Printf("time_stamp      %d\n", b.timeStamp)

	for _, t := range b.transactions {
		t.Print()
	}
}

// Define Blockchain type
type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

// Create blockchain (including genesis block)
func NewBlockchain(blockcainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockcainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

// Create new block
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
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
	return sha256.Sum256([]byte(m))
}

// Marshal block to JSON (translate struct to JSON)
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
		TimeStamp    int64          `json:"time_stamp"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
		TimeStamp:    b.timeStamp,
	})
}

// Add transaction to the transaction pool
func (bc *Blockchain) AddTransaction(
	sender string,
	recipient string,
	value float32,
	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature) bool {

	t := NewTransaction(sender, recipient, value)

	// If sender is mining, add transaction without signature verification
	if sender == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	// Verify transaction signature
	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {

		/*
			// Check if sender has enough funds
			if bc.CalculateTotalAmount(sender) < value {
				log.Println("Error: Not enough balance in sender account")
				return false
			}
		*/

		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("Error: VerifyTransactionSignature failed")

	}
	return false
}

func (bc *Blockchain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, s *utils.Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

// Copy transaction pool
func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)

	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockchainAddress,
				t.recipientBlockchainAddress,
				t.value))
	}

	return transactions
}

// Generate and test Block hash
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)                  // Create a string of zeros
	guessBlock := Block{0, nonce, previousHash, transactions} // Create a new block
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())      // Create a hash from the block

	return guessHashStr[:difficulty] == zeros // Check if the hash starts with the number of zeros
}

// Increment nonce until the hash is valid
func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0

	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}

	return nonce
}

// Mine a block
func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork()             // Find a valid nonce
	previousHash := bc.LastBlock().Hash() // Get the previous hash
	bc.CreateBlock(nonce, previousHash)   // Create a new block with nonce and previous hash
	log.Println("action=mining, status=success")
	return true // Return true if mining is successful
}

// Calculate total amount of a blockchain address
func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0

	// For block in chain
	for _, b := range bc.chain {
		// For transaction in block
		for _, transaction := range b.transactions {
			if transaction.senderBlockchainAddress == blockchainAddress {
				totalAmount -= transaction.value
			}

			if transaction.recipientBlockchainAddress == blockchainAddress {
				totalAmount += transaction.value
			}
		}
	}

	return totalAmount
}

// Create transaction type
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// Create new transaction
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// Print transaction
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address    %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                        %.1f\n", t.value)
}

// Marshal transaction to JSON (translate struct to JSON)
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct { // Creating struct on the fly (only for MarshalJSON)
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
