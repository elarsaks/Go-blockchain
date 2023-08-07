package block

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

// Constants related to blockchain configuration
type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
	port              uint16
	mux               sync.Mutex
	neighbors         []string
	muxNeighbors      sync.Mutex
}

// Create a new instance of Blockchain
func NewBlockchain(blockchainAddress string, port uint16) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	bc.port = port
	return bc
}

// Get chain of the Blockchain
func (bc *Blockchain) Chain() []*Block {
	return bc.chain
}

// Run the Blockchain
func (bc *Blockchain) Run() {
	bc.StartSyncNeighbors()
	bc.ResolveConflicts()
	bc.StartMining() // Start mining automatically
}

// Find neighbors of the Blockchain
func (bc *Blockchain) SetNeighbors() {
	bc.neighbors = utils.FindNeighbors(
		utils.GetHost(), bc.port,
		NEIGHBOR_IP_RANGE_START, NEIGHBOR_IP_RANGE_END,
		BLOCKCHAIN_PORT_RANGE_START, BLOCKCHAIN_PORT_RANGE_END)
	log.Printf("%v", bc.neighbors)
}

// Sync neighbors of the Blockchain
func (bc *Blockchain) SyncNeighbors() {
	bc.muxNeighbors.Lock()
	defer bc.muxNeighbors.Unlock()
	bc.SetNeighbors()
}

// Start syncing neighbors of the Blockchain
func (bc *Blockchain) StartSyncNeighbors() {
	bc.SyncNeighbors()
	_ = time.AfterFunc(time.Second*BLOCKCHIN_NEIGHBOR_SYNC_TIME_SEC, bc.StartSyncNeighbors)
}

// Get the transaction pool the Blockchain
func (bc *Blockchain) TransactionPool() []*Transaction {
	return bc.transactionPool
}

// Empty the transaction pool the Blockchain
func (bc *Blockchain) ClearTransactionPool() {
	bc.transactionPool = bc.transactionPool[:0]
}

// MarshalJSON implements the Marshaler interface for the Blockchain type.
func (bc *Blockchain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"chain"`
	}{
		Blocks: bc.chain,
	})
}

// UnmarshalJSON implements the Unmarshaler interface for the Blockchain type.
func (bc *Blockchain) UnmarshalJSON(data []byte) error {
	v := &struct {
		Blocks *[]*Block `json:"chain"`
	}{
		Blocks: &bc.chain,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}

// Create a new block
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	for _, n := range bc.neighbors {
		endpoint := fmt.Sprintf("http://%s/transactions", n)
		client := &http.Client{}
		req, _ := http.NewRequest("DELETE", endpoint, nil)
		resp, _ := client.Do(req)
		log.Printf("%v", resp)
	}
	return b
}

// Get the last block of the Blockchain
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// Get latest blocks of the Blockchain
func (bc *Blockchain) GetBlocks(amount int) []*Block {
	n := len(bc.chain)
	var last10Blocks []*Block
	if n > amount {
		last10Blocks = append([]*Block(nil), bc.chain[n-amount:n]...)
	} else {
		last10Blocks = append([]*Block(nil), bc.chain...)
	}

	// Reverse the slice
	for i := len(last10Blocks)/2 - 1; i >= 0; i-- {
		opp := len(last10Blocks) - 1 - i
		last10Blocks[i], last10Blocks[opp] = last10Blocks[opp], last10Blocks[i]
	}

	return last10Blocks
}

// Print the Blockchain
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

// Create a new transaction
func (bc *Blockchain) CreateTransaction(sender string, recipient string, message string, value float32,
	senderPublicKey *ecdsa.PublicKey, s *utils.Signature) (bool, error) {

	isTransacted, err := bc.AddTransaction(sender, recipient, message, value, senderPublicKey, s)

	// If there was an error while adding the transaction, log the error and return it
	if err != nil {

		log.Printf("ERROR: %v", err)
		return false, err
	}

	// If the transaction was added successfully, broadcast it to the network
	if isTransacted {
		// Reverse engineer this part of the code
		for _, n := range bc.neighbors {
			publicKeyStr := fmt.Sprintf("%064x%064x", senderPublicKey.X.Bytes(),
				senderPublicKey.Y.Bytes())
			signatureStr := s.String()
			bt := &TransactionRequest{
				&message, &publicKeyStr, &recipient, &sender, &signatureStr, &value}
			m, _ := json.Marshal(bt)
			buf := bytes.NewBuffer(m)
			endpoint := fmt.Sprintf("http://%s/transactions", n)
			client := &http.Client{}
			req, _ := http.NewRequest("PUT", endpoint, buf)
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("ERROR: %v", err)
				return false, err
			}
			log.Printf("%v", resp)
		}
	}

	return isTransacted, nil
}

// AddTransaction adds a new transaction to the transaction pool
// sender is the blockchain address of the sender
// recipient is the blockchain address of the recipient
// message is a text message sent with the transaction
// value is the value of the transaction
// senderPublicKey is the public key of the sender
// s is the signature for the transaction
// The function returns a boolean indicating whether the transaction was added successfully and an error if one occurred
func (bc *Blockchain) AddTransaction(sender string,
	recipient string,
	message string,
	value float32,
	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature) (bool, error) {

	// Create a new transaction
	t := NewTransaction(message, recipient, sender, value)

	// If the sender is the mining address, add the transaction to the pool and return true
	if sender == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true, nil
	}

	// If the transaction signature is not verified, return false and an error
	if !bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		return false, fmt.Errorf("ERROR: Verify Transaction")
	}

	// Calculate the total balance of the sender
	balance, err := bc.CalculateTotalBalance(sender)
	if err != nil {
		// If there is an error calculating the balance, return false and the error
		return false, fmt.Errorf("ERROR: CalculateTotalAmount: %v", err)
	}

	// If the sender's balance is less than the value of the transaction, return false and an error
	if balance < value {
		return false, fmt.Errorf("ERROR: Not enough balance in a wallet")
	}

	// Add the transaction to the transaction pool
	bc.transactionPool = append(bc.transactionPool, t)

	// Return true and no error
	return true, nil
}

// Verify the signature of the transaction
func (bc *Blockchain) VerifyTransactionSignature(

	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature,
	t *Transaction) bool {

	m, _ := json.Marshal(t)

	log.Println("Validate signature", string(m))

	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

// Copy the transaction pool
func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(
				t.senderBlockchainAddress,
				t.recipientBlockchainAddress,
				t.message,
				t.value))
	}
	return transactions
}

// Validate the proof of work
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())

	//* DEBUG #Consensus
	log.Println("VALID PROOF: ", guessHashStr[:difficulty] == zeros)

	return guessHashStr[:difficulty] == zeros
}

// Proof of work
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
// It returns a boolean indicating whether mining was successful.
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

// RegisterNewWallet registers a new wallet on the blockchain
// blockchainAddress is the address of the new wallet
// message is a message that will be associated with the transaction creating the new wallet
// The function returns a boolean indicating whether the wallet was registered successfully
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

// Start mining
func (bc *Blockchain) StartMining() {
	bc.Mining()
	// Schedule the next mining operation to occur after MINING_TIMER_SEC seconds.
	_ = time.AfterFunc(time.Second*MINING_TIMER_SEC, bc.StartMining)
}

//* Calculate the total balance of crypto on the specific address in the Blockchain
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

// Validate the chain
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

// ResolveConflicts resolves conflicts in the blockchain by checking the chains of its neighbors
// and replacing its own chain with the longest valid chain found.
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
