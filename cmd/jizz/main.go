package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/maclovin/jizzjazz/pkg/ascii"
	"github.com/maclovin/jizzjazz/pkg/b64"
	"github.com/maclovin/jizzjazz/pkg/binary"
	"github.com/maclovin/jizzjazz/pkg/gzip"
	"github.com/maclovin/jizzjazz/pkg/hex"
	"github.com/maclovin/jizzjazz/pkg/html"
	"github.com/maclovin/jizzjazz/pkg/octal"
	"github.com/maclovin/jizzjazz/pkg/url"
)

func main() {
	method := flag.String("m", "", "Method of decoding: HTML, ASCII, HEX, Binary, URL, Octal, Base64, GZIP")
	indexRange := flag.String("i", "", "Index range for decoding, e.g., 2:5, :3, 4:")

	flag.Parse()

	if *method == "" {
		fmt.Fprintln(os.Stderr, "Please specify a decoding method")
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
	default:
		return input
	}
}
