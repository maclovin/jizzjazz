package jjReverse

import "strings"

func Encode(input string) string {
	var result strings.Builder

	for i := len(input) - 1; i >= 0; i-- {
		result.WriteByte(input[i])
	}

	return result.String()
}
