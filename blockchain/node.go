package blockchain

import (
	"fmt"
	"sync"
)

// Node represents a node on the blockchain network
type Node struct {
	Address Address
	peers   map[string]bool
	txsPool []*Tx
	txMutex sync.Mutex
	Blockchain
}

// NewNode creates a new node and returns a pointer to it
func NewNode(a Address) *Node {
	bc := NewBlockchain()
	n := Node{Address: a, Blockchain: bc, txsPool: []*Tx{}, peers: make(map[string]bool)}
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

// AddPeer adds a peer to the list
func (n *Node) AddPeer(host string) string {
	n.peers[host] = true
	return host
}

// GetPeers returns the list of peers known to this node
func (n *Node) GetPeers() []string {
	var p []string
	for k := range n.peers {
		p = append(p, k)
	}
	return p
}

func (n *Node) newBlock() *Block {
	t := n.drainTxsPool()
	pb := n.Blockchain.GetLatestBlock()
	b := NewBlock(pb, t)
	return b
}

func (n *Node) addBlock(block *Block) bool {
	ok := block.solve(n.Address)
	if !ok {
		return false
	}
	n.Blockchain.addBlock(block)
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
