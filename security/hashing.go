package security

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Hash(word string) string {
	hash := md5.New()
	io.WriteString(hash, word)

	stringHash := hex.EncodeToString(hash.Sum(nil))

	return stringHash
}
