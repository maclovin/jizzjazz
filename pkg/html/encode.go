package jjHtml

import "html"

func Encode(input string) string {
	return html.EscapeString(input)
}
