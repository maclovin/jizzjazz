package jjSha1

import (
	"crypto/sha1"
	"encoding/hex"
)

func Encode(input string) string {
	hash := sha1.Sum([]byte(input))

	return hex.EncodeToString(hash[:])
}
