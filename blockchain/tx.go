package blockchain

import "time"

// Tx represents a transaction in the blockchain system
type Tx struct {
	NumTxsIn  int      `json:"num_txs_in"`
	TxsIn     []*TxIn  `json:"txs_in"`
	NumTxsOut int      `json:"num_txs_out"`
	TxsOut    []*TxOut `json:"txs_out"`
	Timestamp int64    `json:"timestamp"`
	Hash      Hash     `json:"hash"`
}

// TxIn represents the input of a transaction
type TxIn struct {
	Index      int  `json:"index"`
	PrevOutput Hash `json:"prev_tx_hash"`
}

// TxOut represents the output of a transaction
type TxOut struct {
	Address Address `json:"recipient"`
	Value   int64   `json:"value"`
}

// NewTxIn creates a new transaction input and returns a pointer to it
func NewTxIn(index int, prevOutput Hash) *TxIn {
	t := TxIn{Index: index, PrevOutput: prevOutput}
	return &t
}

// NewTxOut creates a new transaction output and returns a pointer to it
func NewTxOut(recipient Address, value int64) *TxOut {
	t := TxOut{Value: value, Address: recipient}
	return &t
}

// NewTx creates a new transaction from the provided TxIns and TxOuts and returns a pointer to it
func NewTx(txsIn []*TxIn, txsOut []*TxOut) *Tx {
	t := Tx{NumTxsIn: len(txsIn), TxsIn: txsIn, NumTxsOut: len(txsOut), TxsOut: txsOut, Timestamp: time.Now().Unix()}
	t.Hash = calculateHash(t)
	return &t
}

// NewCoinbaseTx creates a new coinbase transaction from the provided address and reward and returns a pointer to it
func NewCoinbaseTx(address Address, reward int64) *Tx {
	var txsOuts []*TxOut
	txsOuts = append(txsOuts, &TxOut{Address: address, Value: reward})
	t := NewTx([]*TxIn{}, txsOuts)
	return t
}
