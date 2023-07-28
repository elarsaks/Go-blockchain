package block

import (
	"encoding/json"
	"fmt"
	"strings"
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
