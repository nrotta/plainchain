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
	ts := time.Now().Unix()
	t := Tx{Sender: sender, Recipient: recipient, Value: value, Timestamp: ts}
	t.Hash = CalculateHash(t)
	return &t
}

// CalculateMerkle returns the Merkle root in SHA256 of the block transactions
// Too keep implementation simple, CalculateMerkle will generate the hash of []*Txs, no a proper merkle tree
func CalculateMerkle(tx []*Tx) Hash {
	h := CalculateHash(tx)
	return h
}
