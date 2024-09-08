package jjB32

import "encoding/base32"

func Encode(input string) string {
	encoded := base32.StdEncoding.EncodeToString([]byte(input))

	return encoded
}
