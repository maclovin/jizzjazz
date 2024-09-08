package jjPunycode

import "golang.org/x/net/idna"

func Encode(input string) string {
	encoded, err := idna.ToASCII(input)

	if err != nil {
		return ""
	}

	return encoded
}
