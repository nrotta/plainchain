# Plainchain - A plain blockchain implementation in golang

Plain chain that implements the following concepts:
- Key blockchain structures (Transactions, Blocks, Blockchain, Node)
- Proof of Work (difficulty target)
- Coinbase reward after successfully mining a Block
- Hashing a Block's transactions (MerkleRoot)
- MemPool for storing trasactions that are yet to be added to a Block

Next steps:
- Support for multiple peer nodes
- Validate, Record, and Query unspent Txs Output (utxos)
