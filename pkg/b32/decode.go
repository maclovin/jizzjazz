package jjB32

import "encoding/base32"

func Decode(input string) string {
	decoded, err := base32.StdEncoding.DecodeString(input)

	if err != nil {
		return ""
	}

	return string(decoded)
}
