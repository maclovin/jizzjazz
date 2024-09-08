package jjKebab

import (
	"strings"
	"unicode"
)

func Encode(input string) string {
	var result strings.Builder

	for i, char := range input {
		if unicode.IsUpper(char) {
			if i > 0 {
				result.WriteString("-")
			}
			result.WriteRune(unicode.ToLower(char))
		} else if char == ' ' {
			result.WriteString("-")
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}
