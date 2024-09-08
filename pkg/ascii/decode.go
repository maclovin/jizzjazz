package jjAscii

import (
	"regexp"
	"strconv"
)

func Decode(input string) string {
	re := regexp.MustCompile(`&#(\d+);`)

	return re.ReplaceAllStringFunc(input, func(s string) string {
		code, _ := strconv.Atoi(re.FindStringSubmatch(s)[1])

		return string(rune(code))
	})
}
