package jjBinary

import (
	"fmt"
	"strings"
)

func Encode(input string) string {
	var sb strings.Builder

	for _, r := range input {
		sb.WriteString(fmt.Sprintf("%08b", r))
	}

	return sb.String()
}
