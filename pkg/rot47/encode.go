package jjRot47

import "strings"

func Encode(input string) string {
	var result strings.Builder
	for _, char := range input {
		if char >= '!' && char <= '~' {
			result.WriteRune(33 + (char-33+47)%94)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
