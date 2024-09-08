package jjHex

import "encoding/hex"

func Decode(input string) string {
	decoded, _ := hex.DecodeString(input)

	return string(decoded)
}
