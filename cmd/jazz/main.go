package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/maclovin/jizzjazz/pkg/ascii"
	"github.com/maclovin/jizzjazz/pkg/binary"
	"github.com/maclovin/jizzjazz/pkg/gzip"
	"github.com/maclovin/jizzjazz/pkg/hex"
	"github.com/maclovin/jizzjazz/pkg/html"
	"github.com/maclovin/jizzjazz/pkg/octal"
	"github.com/maclovin/jizzjazz/pkg/url"
)

func main() {
	method := flag.String("m", "", "Method of encoding: HTML, ASCII, HEX, Binary, URL, Octal and GZIP.")
	indexRange := flag.String("i", "", "Index range for encoding, e.g., 2:5, :3, 4:")

	flag.Parse()

	if *method == "" {
		fmt.Fprintln(os.Stderr, "Please specify an encoding method")
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
		return encode(input, method), nil
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
		return "", fmt.Errorf("Invalid index range")
	}

	return input[:start] + encode(input[start:end], method) + input[end:], nil
}

func encode(input, method string) string {
	switch strings.ToLower(method) {
	case "html":
		return html.Encode(input)
	case "ascii":
		return ascii.Encode(input)
	case "hex":
		return hex.Encode(input)
	case "binary":
		return binary.Encode(input)
	case "url":
		return url.Encode(input)
	case "octal":
		return octal.Encode(input)
	case "base64":
		return base64.Encode(input)
	case "gzip":
		return gzip.Encode(input)
	case "md5":
		return md5.Encode(input)
	case "sha256":
		return sha256.Encode(input)
	default:
		return input
	}
}
