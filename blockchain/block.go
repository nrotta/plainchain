package blockchain

import (
	"bytes"
	"math"
	"time"
)

// Header represents the header of a block
type Header struct {
	Nonce      uint32 `json:"nonce"`
	PrevHash   Hash   `json:"prev_hash"`
	MerkleRoot Hash   `json:"merkle_root"`
	Timestamp  int64  `json:"timestamp"`
	NumTxs     int    `json:"num_txs"`
}

// Block represents a block in the blockchain
type Block struct {
	Header `json:"header"`
	Height uint32 `json:"height"`
	Hash   Hash   `json:"hash"`
	Txs    []*Tx  `json:"txs"`
}

// NewBlock creates a new Block with the specified transactions, links it to the previous block on the chain and returns a pointer to it
func NewBlock(prevBlock *Block, Txs *[]*Tx) *Block {
	h := prevBlock.Height + 1
	t := time.Now().Unix()
	b := Block{Height: h, Header: Header{PrevHash: prevBlock.Hash, Timestamp: t, NumTxs: len(*Txs)}, Txs: *Txs}
	return &b
}

// Solve adds the coinbase Tx, calculates the Merkle root of the Txs, and then generates the block hash by iterating the nonce
func (b *Block) solve(address Address, difficulty int) bool {
	b.addCoinbaseTx(address)
	b.MerkleRoot = calculateMerkle(b.Txs)
	target := make([]byte, difficulty)

	for b.Nonce = uint32(0); b.Nonce <= math.MaxUint32; b.Nonce++ {
		b.Hash = calculateHash(calculateHash(b.Header))
		if bytes.Compare(target, b.Hash[:difficulty]) == 0 {
			return true
		}
	}
	return false
}

func (b *Block) addCoinbaseTx(address Address) {
	r := CalculateBlockReward(b.Height)
	t := NewTx(Address{}, address, r)
	b.Txs = append([]*Tx{t}, b.Txs...)
	b.NumTxs++
}
