package jjB45

import "strings"

const base45Charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"

func Encode(input string) string {
	var result strings.Builder
	data := []byte(input)

	for i := 0; i < len(data); i += 2 {
		if i+1 < len(data) {
			value := int(data[i])<<8 + int(data[i+1])
			result.WriteByte(base45Charset[value%45])
			result.WriteByte(base45Charset[(value/45)%45])
			result.WriteByte(base45Charset[(value/2025)%45])
		} else {
			value := int(data[i])
			result.WriteByte(base45Charset[value%45])
			result.WriteByte(base45Charset[(value/45)%45])
		}
	}

	return result.String()
}
