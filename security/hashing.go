package security

import (
	"crypto/sha256"
	"io"
)

func Hash(word string) ([]byte, error) {
	h256 := sha256.New()
	io.WriteString(h256, "Hello money money!")

	return h256.Sum(nil), nil
}
