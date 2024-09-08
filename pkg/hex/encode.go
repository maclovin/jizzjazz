package jjHex

import "encoding/hex"

func Encode(input string) string {
	return hex.EncodeToString([]byte(input))
}
