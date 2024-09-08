package jjB45

import "strings"

func Decode(input string) string {
	var result strings.Builder
	data := []rune(input)
	length := len(data)

	for i := 0; i < length; i += 3 {
		if i+2 < length {
			value := strings.IndexRune(base45Charset, data[i]) + 45*strings.IndexRune(base45Charset, data[i+1]) + 2025*strings.IndexRune(base45Charset, data[i+2])
			result.WriteByte(byte(value >> 8))
			result.WriteByte(byte(value & 0xFF))
		} else if i+1 < length {
			value := strings.IndexRune(base45Charset, data[i]) + 45*strings.IndexRune(base45Charset, data[i+1])
			result.WriteByte(byte(value))
		} else {
			return ""
		}
	}

	return result.String()
}
