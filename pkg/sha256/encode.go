package jjSha256

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encode(input string) string {
	hash := sha256.Sum256([]byte(input))

	return hex.EncodeToString(hash[:])
}
