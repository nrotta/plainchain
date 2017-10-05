package blockchain

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// Node represents a node on the blockchain p2p network
type Node struct {
	UUID       uuid.UUID
	Address    Address
	Difficulty int
	TxsPool    []*Tx
	txMutex    sync.Mutex
	Blockchain
}

// NewNode creates a new p2p node and returns a pointer to it
func NewNode(address string, difficulty int) *Node {
	bc := NewBlockchain()
	n := Node{UUID: uuid.New(), Address: Address(address), Difficulty: difficulty, Blockchain: bc, TxsPool: []*Tx{}}
	return &n
}

// Mine runs the mining process until the node is shutted down
func (n *Node) Mine() {
	for {
		b := n.NewBlock()
		ok := n.AddBlock(b)
		if !ok {
			fmt.Printf("Mining was unsuccessful for Block %d\n", b.Height)
			continue
		}

		fmt.Printf("Block mined and added to the chain: [Height: %d, Nonce: %d, Hash: %x, PrevHash: %x, NumTxs: %d]\n", b.Height, b.Nonce, b.Hash, b.PrevHash, b.NumTxs)
	}
}

// AddTx adds a transaction to the transaction pool
func (n *Node) AddTx(tx *Tx) {
	n.txMutex.Lock()
	n.TxsPool = append(n.TxsPool, tx)
	n.txMutex.Unlock()
}

// NewBlock creates a new Block using the transactions from the transactions pool and returns a pointer to it
func (n *Node) NewBlock() *Block {
	t := n.drainTxsPool()
	pb := n.Blockchain.GetLatestBlock()
	b := NewBlock(pb, t)
	return b
}

// AddBlock mines a given block and adds it to the blockchain
func (n *Node) AddBlock(block *Block) bool {
	ok := block.Solve(n.Address, n.Difficulty)
	if !ok {
		return false
	}
	n.Blockchain.AddBlock(block)
	return true
}

func (n *Node) drainTxsPool() *[]*Tx {
	n.txMutex.Lock()
	t := make([]*Tx, len(n.TxsPool))
	copy(t, n.TxsPool)
	n.TxsPool = []*Tx{}
	n.txMutex.Unlock()
	return &t
}
