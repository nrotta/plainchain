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
	txsPool    []*Tx
	txMutex    sync.Mutex
	Blockchain
}

// NewNode creates a new p2p node and returns a pointer to it
func NewNode(address string, difficulty int) *Node {
	bc := NewBlockchain()
	n := Node{UUID: uuid.New(), Address: Address(address), Difficulty: difficulty, Blockchain: bc, txsPool: []*Tx{}}
	return &n
}

// GetTxsPool returns the Txs in the pool for this node
func (n *Node) GetTxsPool() []*Tx {
	return n.txsPool
}

// Mine runs the mining process until the node is shutted down
func (n *Node) Mine() {
	for {
		b := n.newBlock()
		ok := n.addBlock(b)
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
	n.txsPool = append(n.txsPool, tx)
	n.txMutex.Unlock()
}

func (n *Node) newBlock() *Block {
	t := n.drainTxsPool()
	pb := n.Blockchain.GetLatestBlock()
	b := NewBlock(pb, t)
	return b
}

func (n *Node) addBlock(block *Block) bool {
	ok := block.Solve(n.Address, n.Difficulty)
	if !ok {
		return false
	}
	n.Blockchain.AddBlock(block)
	return true
}

func (n *Node) drainTxsPool() *[]*Tx {
	n.txMutex.Lock()
	t := make([]*Tx, len(n.txsPool))
	copy(t, n.txsPool)
	n.txsPool = []*Tx{}
	n.txMutex.Unlock()
	return &t
}
