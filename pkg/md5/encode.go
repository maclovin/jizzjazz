package jjMd5

import (
	"crypto/md5"
	"encoding/hex"
)

func Encode(input string) string {
	hash := md5.Sum([]byte(input))

	return hex.EncodeToString(hash[:])
}
