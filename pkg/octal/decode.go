package jjOctal

import (
	"fmt"
	"strings"
)

func Decode(input string) string {
	var sb strings.Builder

	for i := 0; i < len(input); i += 4 {
		var b byte

		fmt.Sscanf(input[i:i+4], "\\%o", &b)
		sb.WriteByte(b)
	}

	return sb.String()
}
