package blockchain

import (
	"crypto/sha256"
	"fmt"
)

// Hash is a alias for [32]byte used for hashes
type Hash [32]byte

// MarshalText (h Hash) returns the hash value as a hex string used for json marshalling
func (h Hash) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%x", h)), nil
}

// Address is a alias for []byte used for blockchain addresses
type Address []byte

// MarshalText (a Address) returns the address as a hex string used for json marshalling
func (a Address) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", a)), nil
}

// CalculateHash returns the SHA256 digest of the interface obj
func CalculateHash(obj interface{}) Hash {
	d := fmt.Sprintf("%v", obj)
	h := sha256.Sum256([]byte(d))
	return h
}

// CalculateBlockReward calculates the reward that must be granted for the block being mined
func CalculateBlockReward(height uint32) int64 {
	return 5000000000 >> uint(height/210000) //int64()
}
