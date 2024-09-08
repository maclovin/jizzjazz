package jjUrl

import "net/url"

func Encode(input string) string {
	return url.QueryEscape(input)
}
