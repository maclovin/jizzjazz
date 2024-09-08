package jjUnicode

import (
	"regexp"
	"strconv"
	"strings"
)

func Decode(input string) string {
	re := regexp.MustCompile(`\\u([0-9A-Fa-f]{4})`)
	matches := re.FindAllStringSubmatch(input, -1)

	var result strings.Builder
	lastIndex := 0
	for _, match := range matches {
		start := strings.Index(input[lastIndex:], match[0])
		result.WriteString(input[lastIndex : lastIndex+start])
		lastIndex += start + len(match[0])

		codePoint, err := strconv.ParseInt(match[1], 16, 32)
		if err != nil {
			return ""
		}
		result.WriteRune(rune(codePoint))
	}

	result.WriteString(input[lastIndex:])

	return result.String()
}
