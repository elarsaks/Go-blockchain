package block

import (
	"encoding/json"
	"fmt"
	"strings"
)

// --- Types ---
// Transaction represents a single transaction in the blockchain.
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// TransactionRequest represents a request to create a new transaction.
type TransactionRequest struct {
	SenderBlockchainAddress    *string  `json:"senderBlockchainAddress"`
	RecipientBlockchainAddress *string  `json:"recipientBlockchainAddress"`
	SenderPublicKey            *string  `json:"senderPublicKey"`
	Value                      *float32 `json:"value"`
	Signature                  *string  `json:"signature"`
}

// AmountResponse represents the response with the amount in a transaction.
type AmountResponse struct {
	Amount float32 `json:"amount"`
}

// --- Functions ---
// NewTransaction creates a new transaction.
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// --- Methods ---
// Print outputs the details of the transaction.
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" senderBlockchainAddress      %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipientBlockchainAddress   %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                          %.1f\n", t.value)
}

// MarshalJSON implements the Marshaler interface for the Transaction type.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"senderBlockchainAddress"`
		Recipient string  `json:"recipientBlockchainAddress"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

// UnmarshalJSON implements the Unmarshaler interface for the Transaction type.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	v := &struct {
		Sender    *string  `json:"senderBlockchainAddress"`
		Recipient *string  `json:"recipientBlockchainAddress"`
		Value     *float32 `json:"value"`
	}{
		Sender:    &t.senderBlockchainAddress,
		Recipient: &t.recipientBlockchainAddress,
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
		tr.Value == nil ||
		tr.Signature == nil {
		return false
	}
	return true
}

// MarshalJSON implements the Marshaler interface for the AmountResponse type.
func (ar *AmountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Amount float32 `json:"amount"`
	}{
		Amount: ar.Amount,
	})
}
