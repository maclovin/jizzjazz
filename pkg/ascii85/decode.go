package jjAscii85

import "encoding/ascii85"

func Decode(input string) string {
	decoded := make([]byte, len(input))
	n, _, err := ascii85.Decode(decoded, []byte(input), true)

	if err != nil {
		return ""
	}

	return string(decoded[:n])
}
