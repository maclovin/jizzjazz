package jjBinary

import (
	"strconv"
	"strings"
)

func Decode(input string) string {
	var sb strings.Builder

	for i := 0; i < len(input); i += 8 {
		byteStr := input[i : i+8]
		byteVal, _ := strconv.ParseUint(byteStr, 2, 8)
		sb.WriteByte(byte(byteVal))
	}

	return sb.String()
}
