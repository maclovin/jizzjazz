package jjHtml

import "html"

func Decode(input string) string {
	return html.UnescapeString(input)
}
