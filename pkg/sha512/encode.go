package jjSha512

import (
	"crypto/sha512"
	"encoding/hex"
)

func Encode(input string) string {
	hash := sha512.Sum512([]byte(input))

	return hex.EncodeToString(hash[:])
}
