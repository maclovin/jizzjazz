package jjCamel

import "strings"

func Encode(input string) string {
	var result strings.Builder
	words := strings.Fields(input)

	for i, word := range words {
		if i == 0 {
			result.WriteString(strings.ToLower(word))
		} else {
			result.WriteString(strings.Title(word))
		}
	}

	return result.String()
}
