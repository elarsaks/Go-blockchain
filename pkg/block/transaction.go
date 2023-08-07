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

	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

// --- Types ---
// Transaction represents a single transaction in the blockchain.
type Transaction struct {
	message                    string
	recipientBlockchainAddress string
	senderBlockchainAddress    string
	value                      float32
}

// TransactionRequest represents a request to create a new transaction.
type TransactionRequest struct {
	Message                    *string  `json:"message"`
	RecipientBlockchainAddress *string  `json:"recipientBlockchainAddress"`
	SenderBlockchainAddress    *string  `json:"senderBlockchainAddress"`
	SenderPublicKey            *string  `json:"senderPublicKey"`
	Signature                  *string  `json:"signature"`
	Value                      *float32 `json:"value"`
}

// AmountResponse represents the response with the amount in a transaction.
type BalanceResponse struct {
	Balance float32 `json:"balance"`
	Error   string  `json:"error"`
}

// --- Functions ---
// NewTransaction creates a new transaction.
func NewTransaction(sender string, recipient string, message string, value float32) *Transaction {
	return &Transaction{sender, recipient, message, value}
}

// --- Methods ---
// Print outputs the details of the transaction.
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" senderBlockchainAddress      %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipientBlockchainAddress   %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" message                      %s\n", t.message)
	fmt.Printf(" value                          %.1f\n", t.value)
}

// MarshalJSON implements the Marshaler interface for the Transaction type.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Message   string  `json:"message"`
		Recipient string  `json:"recipientBlockchainAddress"`
		Sender    string  `json:"senderBlockchainAddress"`
		Value     float32 `json:"value"`
	}{
		Message:   t.message,
		Recipient: t.recipientBlockchainAddress,
		Sender:    t.senderBlockchainAddress,
		Value:     t.value,
	})
}

// UnmarshalJSON implements the Unmarshaler interface for the Transaction type.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	v := &struct {
		Message   *string  `json:"message"`
		Recipient *string  `json:"recipientBlockchainAddress"`
		Sender    *string  `json:"senderBlockchainAddress"`
		Value     *float32 `json:"value"`
	}{
		Message:   &t.message,
		Recipient: &t.recipientBlockchainAddress,
		Sender:    &t.senderBlockchainAddress,
		Value:     &t.value,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}

// Validate checks if the transaction request is valid.
func (tr *TransactionRequest) Validate() bool {
	if tr.SenderBlockchainAddress == nil ||
		tr.RecipientBlockchainAddress == nil ||
		tr.SenderPublicKey == nil ||
		tr.Message == nil ||
		tr.Value == nil ||
		tr.Signature == nil {
		return false
	}
	return true
}

// MarshalJSON implements the Marshaler interface for the AmountResponse type.
func (br *BalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Balance float32 `json:"balance"`
		Error   string  `json:"error"`
	}{
		Balance: br.Balance,
		Error:   br.Error,
	})
}

// Get the transaction pool the Blockchain
func (bc *Blockchain) TransactionPool() []*Transaction {
	return bc.transactionPool
}

// Empty the transaction pool the Blockchain
func (bc *Blockchain) ClearTransactionPool() {
	bc.transactionPool = bc.transactionPool[:0]
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
