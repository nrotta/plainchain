package blockchain

// Blockchain represents the current blockchain
type Blockchain struct {
	height      uint32
	chain       map[Hash]*Block
	blocksIndex map[uint32]*Block
	txsIndex    map[Hash]*Tx
}

func newBlockchain() Blockchain {
	bc := Blockchain{chain: make(map[Hash]*Block), blocksIndex: make(map[uint32]*Block), txsIndex: make(map[Hash]*Tx)}
	return bc
}

// GetLatestBlock returns a pointer to the latest Block in the chain or a nil Block if the chain is empty
func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.GetBlock(bc.height)
}

// GetBlock returns a pointer to the Block located at height defined as param
func (bc *Blockchain) GetBlock(height uint32) *Block {
	b := &Block{}
	if height > 0 {
		b = bc.blocksIndex[height]
	}
	return b
}

// GetTx returns a pointer to the Tx with the Hash defined as param
func (bc *Blockchain) GetTx(hash Hash) *Tx {
	return bc.txsIndex[hash]
}

func (bc *Blockchain) addBlock(block *Block) {
	bc.chain[block.Hash] = block
	bc.blocksIndex[block.Height] = block
	bc.height++

	for _, t := range block.Txs {
		bc.txsIndex[t.Hash] = t
	}
}
