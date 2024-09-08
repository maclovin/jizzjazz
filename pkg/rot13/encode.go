package jjRot13

import "strings"

func Encode(input string) string {
	var result strings.Builder

	for _, char := range input {
		switch {
		case 'a' <= char && char <= 'z':
			result.WriteRune('a' + (char-'a'+13)%26)
		case 'A' <= char && char <= 'Z':
			result.WriteRune('A' + (char-'A'+13)%26)
		default:
			result.WriteRune(char)
		}
	}

	return result.String()
}
