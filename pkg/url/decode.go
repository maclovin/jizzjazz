package jjUrl

import "net/url"

func Decode(input string) string {
	decoded, _ := url.QueryUnescape(input)

	return decoded
}
