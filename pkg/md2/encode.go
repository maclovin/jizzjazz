package jjMd2

import (
	"encoding/hex"

	"github.com/htruong/go-md2"
)

func Encode(input string) string {
	hash := md2.New()
	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))
}
