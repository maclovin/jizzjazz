package jjPunycode

import "golang.org/x/net/idna"

func Decode(input string) string {
	decoded, err := idna.ToUnicode(input)

	if err != nil {
		return ""
	}

	return decoded
}
