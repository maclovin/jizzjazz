package jjSha384

import (
	"crypto/sha512"
	"encoding/hex"
)

func Encode(input string) string {
	hash := sha512.Sum384([]byte(input))

	return hex.EncodeToString(hash[:])
}
