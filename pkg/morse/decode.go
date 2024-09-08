package jjMorse

import "strings"

func Decode(input string) string {
	reverseMorseCode := make(map[string]rune)

	for k, v := range morseCode {
		reverseMorseCode[v] = k
	}

	var result strings.Builder

	for _, word := range strings.Split(input, " / ") {
		for _, letter := range strings.Split(word, " ") {
			if char, ok := reverseMorseCode[letter]; ok {
				result.WriteRune(char)
			}
		}
		result.WriteRune(' ')
	}

	return strings.TrimSpace(result.String())
}
