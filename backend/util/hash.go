package util

import (
	"crypto/sha256"
	"math/big"
)

func StringToUint248Hash(input string) *big.Int {
	// Hash the input string using SHA-256
	hash := sha256.Sum256([]byte(input))

	// Extract the first 31 bytes (248 bits) from the hash
	hashBytes := hash[:31]

	// Convert the extracted bytes to a big.Int
	uint248Value := new(big.Int).SetBytes(hashBytes)

	return uint248Value
}
