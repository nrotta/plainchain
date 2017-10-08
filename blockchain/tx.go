package blockchain

import "time"

// Tx represents a transaction in the blockchain system
type Tx struct {
	Sender    Address `json:"sender"`
	Recipient Address `json:"recipient"`
	Value     int64   `json:"value"`
	Timestamp int64   `json:"timestamp"`
	Hash      Hash    `json:"hash"`
}

// NewTx creates a new transaction returns a pointer to it
func NewTx(sender, recipient Address, value int64) *Tx {
	t := Tx{Sender: sender, Recipient: recipient, Value: value, Timestamp: time.Now().Unix()}
	t.Hash = calculateHash(t)
	return &t
}
