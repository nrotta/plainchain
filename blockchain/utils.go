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

func calculateHash(obj interface{}) Hash {
	d := fmt.Sprintf("%v", obj)
	h := sha256.Sum256([]byte(d))
	return h
}
