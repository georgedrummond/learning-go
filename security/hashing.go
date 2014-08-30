package security

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Hash(word string) string {
	h256 := sha256.New()
	io.WriteString(h256, "This is my string to hash")

	stringHash := hex.EncodeToString(h256.Sum(nil))

	return stringHash
}
