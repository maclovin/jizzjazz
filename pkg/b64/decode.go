package jjB64

import "encoding/base64"

func Decode(input string) string {
	decoded, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		return ""
	}

	return string(decoded)
}
