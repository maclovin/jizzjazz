package jjUnicode

import (
	"fmt"
	"strings"
)

func Encode(input string) string {
	var result strings.Builder

	for _, char := range input {
		result.WriteString(fmt.Sprintf("\\u%04X", char))
	}

	return result.String()
}
