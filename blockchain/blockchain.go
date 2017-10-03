package blockchain

// Blockchain represents the current blockchain
type Blockchain struct {
	Height      uint32
	Chain       map[Hash]*Block
	BlocksIndex map[uint32]*Block
	TxsIndex    map[Hash]*Tx
}

// NewBlockchain creates a new chain and returns it
func NewBlockchain() Blockchain {
	bc := Blockchain{Chain: make(map[Hash]*Block), BlocksIndex: make(map[uint32]*Block), TxsIndex: make(map[Hash]*Tx)}
	return bc
}

// GetLatestBlock returns a pointer to the latest Block in the chain or a nil Block if the chain is empty
func (bc *Blockchain) GetLatestBlock() *Block {
	b := &Block{}
	if bc.Height > 0 {
		b = bc.BlocksIndex[bc.Height]
	}
	return b
}

// AddBlock adds a given block to the blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	bc.Chain[block.Hash] = block
	bc.BlocksIndex[block.Height] = block
	bc.AddTxsToIndex(block.Txs)
	bc.Height++
}

// AddTxsToIndex adds the given transactions to the transactions index
func (bc *Blockchain) AddTxsToIndex(txs []*Tx) {
	for _, t := range txs {
		bc.TxsIndex[t.Hash] = t
	}
}
