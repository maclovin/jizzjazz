package jjAscii85

import "encoding/ascii85"

func Encode(input string) string {
	byteData := []byte(input)
	encoded := make([]byte, ascii85.MaxEncodedLen(len(byteData)))

	ascii85.Encode(encoded, byteData)

	return string(encoded[:ascii85.Encode(encoded, byteData)])
}
