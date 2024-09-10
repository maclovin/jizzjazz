package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/maclovin/jizzjazz/cmd/shared"
	"github.com/maclovin/jizzjazz/pkg/ascii"
	"github.com/maclovin/jizzjazz/pkg/ascii85"
	"github.com/maclovin/jizzjazz/pkg/atbash"
	"github.com/maclovin/jizzjazz/pkg/b32"
	"github.com/maclovin/jizzjazz/pkg/b45"
	"github.com/maclovin/jizzjazz/pkg/b64"
	"github.com/maclovin/jizzjazz/pkg/binary"
	"github.com/maclovin/jizzjazz/pkg/gzip"
	"github.com/maclovin/jizzjazz/pkg/hex"
	"github.com/maclovin/jizzjazz/pkg/html"
	"github.com/maclovin/jizzjazz/pkg/morse"
	"github.com/maclovin/jizzjazz/pkg/octal"
	"github.com/maclovin/jizzjazz/pkg/punycode"
	"github.com/maclovin/jizzjazz/pkg/unicode"
	"github.com/maclovin/jizzjazz/pkg/url"
)

func main() {
	method := flag.String("m", "", "Allowed Reverse Encoding/Transformation Methods: html, url, ascii, hex, gzip, unicode, etc... (-h to see the full methods list)")
	indexRange := flag.String("i", "", "Index range for Encode/Transform, e.g., 2:5, :3, 4:")
	help := flag.Bool("h", false, "Show all the methods and usage")

	flag.Parse()

	if *method == "" || *help {
		fmt.Printf("%sUsage%s\n$ echo \"YOUR STRING CONTENT\" | ./jizz -m [METHOD]\n", shared.BoldFont, shared.ResetFont)
		fmt.Printf("\n%sExample%s\n$ echo \"48656c6c6f20576f726c640a\" | ./jizz -m hex\n$ Hello World\n", shared.BoldFont, shared.ResetFont)

		showMethodsList()

		return
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
		return
	}

	processedInput, err := processInputByRange(string(input), *indexRange, *method)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error processing input:", err)
		return
	}

	fmt.Print(processedInput)
}

func showMethodsList() {
	shared.PrintTable(
		[]string{"Method", "Type"},
		[][]string{
			{"html", "Encoding"},
			{"ascii", "Encoding"},
			{"hex", "Encoding"},
			{"binary", "Encoding"},
			{"url", "Encoding"},
			{"octal", "Encoding"},
			{"base32", "Encoding"},
			{"base45", "Encoding"},
			{"base64", "Encoding"},
			{"gzip", "Encoding"},
			{"ascii85", "Encoding"},
			{"atbash", "Transformation"},
			{"morse", "Encoding"},
			{"punycode", "Encoding"},
			{"unicode", "Encoding"},
		},
	)
}

func processInputByRange(input, indexRange, method string) (string, error) {
	if indexRange == "" {
		return decode(input, method), nil
	}

	parts := strings.Split(indexRange, ":")
	start, end := 0, len(input)

	if parts[0] != "" {
		start, _ = strconv.Atoi(parts[0])
	}
	if len(parts) > 1 && parts[1] != "" {
		end, _ = strconv.Atoi(parts[1])
	}

	if start > end || start < 0 || end > len(input) {
		return "", fmt.Errorf("invalid index range")
	}

	return input[:start] + decode(input[start:end], method) + input[end:], nil
}

func decode(input, method string) string {
	switch strings.ToLower(method) {
	case "html":
		return jjHtml.Decode(input)
	case "ascii":
		return jjAscii.Decode(input)
	case "hex":
		return jjHex.Decode(input)
	case "binary":
		return jjBinary.Decode(input)
	case "url":
		return jjUrl.Decode(input)
	case "octal":
		return jjOctal.Decode(input)
	case "base64":
		return jjB64.Decode(input)
	case "gzip":
		return jjGzip.Decode(input)
	case "ascii85":
		return jjAscii85.Decode(input)
	case "atbash":
		return jjAtbash.Decode(input)
	case "base32":
		return jjB32.Decode(input)
	case "base45":
		return jjB45.Decode(input)
	case "morse":
		return jjMorse.Decode(input)
	case "punycode":
		return jjPunycode.Decode(input)
	case "unicode":
		return jjUnicode.Decode(input)

	default:
		showMethodsList()
		return input
	}
}
