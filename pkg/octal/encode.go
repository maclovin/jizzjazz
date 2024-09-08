package jjOctal

import (
	"fmt"
	"strings"
)

func Encode(input string) string {
	var sb strings.Builder

	for _, r := range input {
		sb.WriteString(fmt.Sprintf("\\%o", r))
	}

	return sb.String()
}
