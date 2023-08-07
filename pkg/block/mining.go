package block

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// ==============================
// Blockchain Proof and Mining Methods
// ==============================

// ValidProof validates the proof of work.
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())

	//* DEBUG #Consensus
	log.Println("VALID PROOF: ", guessHashStr[:difficulty] == zeros)

	return guessHashStr[:difficulty] == zeros
}

// ProofOfWork finds the proof of work.
func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

// Mining creates a new block and adds it to the blockchain.
func (bc *Blockchain) Mining() bool {
	// Lock the blockchain while mining
	bc.mux.Lock()
	defer bc.mux.Unlock()

	// Log out blockchain
	// bc.Print() // TODO: Remove debug

	//* DEBUG #Consensus Wallet registration mining should be done some where else
	// Don't mine when there is no transaction and blockchain already has few blocks
	if len(bc.transactionPool) == 0 {
		return false
	}

	// Add a mining reward transaction
	_, err := bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, "MINING REWARD", MINING_REWARD, nil, nil)

	// If an error occurred adding the transaction, log the error and return false
	if err != nil {
		log.Printf("ERROR: %v", err)
		return false
	}

	// Find a new proof of work and create a new block
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)

	// Log a successful mining operation
	// #debug
	log.Println("action=mining, status=success")

	// Send a consensus request to each neighbor
	for _, n := range bc.neighbors {

		fmt.Println("Send consensus to neigbour ", n)

		endpoint := fmt.Sprintf("http://%s/consensus", n)
		client := &http.Client{}
		req, _ := http.NewRequest("PUT", endpoint, nil)
		resp, err := client.Do(req)

		// If an error occurred making the request, log the error
		if err != nil {
			log.Printf("ERROR: %v", err)
			return false
		}

		log.Printf("%v", resp)
	}

	// Return true indicating the mining operation was successful
	return true
}

// StartMining initiates the mining process.
func (bc *Blockchain) StartMining() {
	bc.Mining()
	// Schedule the next mining operation to occur after MINING_TIMER_SEC seconds.
	_ = time.AfterFunc(time.Second*MINING_TIMER_SEC, bc.StartMining)
}

// ==============================
// Blockchain Wallet and Balance Methods
// ==============================

// RegisterNewWallet registers a new wallet on the blockchain.
func (bc *Blockchain) RegisterNewWallet(blockchainAddress string, message string) bool {

	// Add a transaction for the new wallet
	_, err := bc.AddTransaction(MINING_SENDER, blockchainAddress, message, 0, nil, nil)

	// If an error occurred adding the transaction, log the error and return false
	if err != nil {
		log.Printf("ERROR: %v", err)
		return false
	}

	// Mine a new block when the wallet is registered successfully
	bc.StartMining()

	// Return true indicating the wallet was registered successfully
	return true
}

// CalculateTotalBalance calculates the total balance of crypto on the specific address in the Blockchain.
func (bc *Blockchain) CalculateTotalBalance(blockchainAddress string) (float32, error) {
	var totalBalance float32 = 0.0
	addressFound := false

	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value

			if blockchainAddress == t.recipientBlockchainAddress {
				totalBalance += value
				addressFound = true
			}

			if blockchainAddress == t.senderBlockchainAddress {
				totalBalance -= value
				addressFound = true
			}
		}
	}

	if !addressFound {
		return 0.0, fmt.Errorf("Address not found in the Blockchain")
	}

	return totalBalance, nil
}

// ==============================
// Blockchain Chain Validation and Conflict Resolution Methods
// ==============================

// ValidChain validates the chain.
func (bc *Blockchain) ValidChain(chain []*Block) bool {
	preBlock := chain[0]
	currentIndex := 1
	for currentIndex < len(chain) {
		b := chain[currentIndex]
		if b.previousHash != preBlock.Hash() {
			return false
		}

		if !bc.ValidProof(b.Nonce(), b.PreviousHash(), b.Transactions(), MINING_DIFFICULTY) {
			return false
		}

		preBlock = b
		currentIndex += 1
	}
	return true
}

// ResolveConflicts resolves conflicts in the blockchain by checking the chains of its neighbors.
func (bc *Blockchain) ResolveConflicts() bool {
	// Initialize variables to track the longest chain and its length
	var longestChain []*Block = nil
	maxLength := len(bc.chain)

	// Iterate over the neighbors to fetch their chains
	for _, n := range bc.neighbors {
		fmt.Println("Resolve conflict with ", n)

		// Construct the endpoint URL to fetch the chain from the neighbor
		endpoint := fmt.Sprintf("http://%s/chain", n)

		// Send an HTTP GET request to the neighbor's endpoint to fetch their chain
		resp, err := http.Get(endpoint)
		if err != nil {

			// Log any error that occurred while fetching the chain
			log.Printf("ERROR: Failed to fetch chain from neighbor %s: %v", n, err)
			continue // Skip to the next neighbor in case of error
		}

		// Check the response status code to see if the request was successful
		if resp.StatusCode == http.StatusOK {
			var bcResp Blockchain
			decoder := json.NewDecoder(resp.Body)

			// Decode the JSON response into a Blockchain object
			err := decoder.Decode(&bcResp)
			if err != nil {
				// Log any error that occurred during JSON decoding
				log.Printf("ERROR: Failed to decode JSON response from neighbor %s: %v", n, err)
				continue // Skip to the next neighbor in case of error
			}

			// Get the chain from the neighbor's Blockchain object
			chain := bcResp.Chain()

			// Check if the fetched chain is longer than the current longest chain
			// and if it is a valid chain using bc.ValidChain()
			if len(chain) > maxLength && bc.ValidChain(chain) {
				maxLength = len(chain)
				longestChain = chain
			}
		} else {
			// Log the status code if the request to the neighbor's endpoint was not successful
			log.Printf("WARNING: Failed to fetch chain from neighbor %s. Status code: %d", n, resp.StatusCode)
		}
	}

	// If a longer valid chain was found, replace the blockchain's chain with it
	if longestChain != nil {
		bc.chain = longestChain
		log.Printf("INFO: Resolved conflicts. Replaced blockchain with the longest valid chain.")
		return true
	}

	// If no longer valid chain was found, log and return false
	log.Printf("INFO: No longer valid chain found among neighbors. No conflicts resolved.")
	return false
}
