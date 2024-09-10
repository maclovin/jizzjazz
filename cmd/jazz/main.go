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
	"github.com/maclovin/jizzjazz/pkg/camel"
	"github.com/maclovin/jizzjazz/pkg/crc32"
	"github.com/maclovin/jizzjazz/pkg/gzip"
	"github.com/maclovin/jizzjazz/pkg/hex"
	"github.com/maclovin/jizzjazz/pkg/html"
	"github.com/maclovin/jizzjazz/pkg/kebab"
	"github.com/maclovin/jizzjazz/pkg/lower"
	"github.com/maclovin/jizzjazz/pkg/md2"
	"github.com/maclovin/jizzjazz/pkg/md5"
	"github.com/maclovin/jizzjazz/pkg/morse"
	"github.com/maclovin/jizzjazz/pkg/octal"
	"github.com/maclovin/jizzjazz/pkg/punycode"
	"github.com/maclovin/jizzjazz/pkg/reverse"
	"github.com/maclovin/jizzjazz/pkg/rot13"
	"github.com/maclovin/jizzjazz/pkg/rot18"
	"github.com/maclovin/jizzjazz/pkg/rot47"
	"github.com/maclovin/jizzjazz/pkg/sha1"
	"github.com/maclovin/jizzjazz/pkg/sha256"
	"github.com/maclovin/jizzjazz/pkg/sha384"
	"github.com/maclovin/jizzjazz/pkg/sha512"
	"github.com/maclovin/jizzjazz/pkg/unicode"
	"github.com/maclovin/jizzjazz/pkg/upper"
	"github.com/maclovin/jizzjazz/pkg/url"
)

func main() {
	method := flag.String("m", "", "Allowed Hashing/Encoding/Cyphering/Transformation Methods: html, ascii, hex, binary, gzip, sha256, etc... (-h to see the full methods list)")
	indexRange := flag.String("i", "", "Index range for encoding, e.g., 2:5, :3, 4:")
	help := flag.Bool("h", false, "Show all the methods and usage")

	flag.Parse()

	if *method == "" || *help {
		fmt.Printf("%sUsage%s\n$ echo \"YOUR STRING CONTENT\" | ./jazz -m [METHOD]\n", shared.BoldFont, shared.ResetFont)
		fmt.Printf("\n%sExample%s\n$ echo \"Hello World\" | ./jazz -m rot18\n$ Uryyb Jbeyq\n", shared.BoldFont, shared.ResetFont)

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
			{"md5", "Hashing"},
			{"sha256", "Hashing"},
			{"ascii85", "Encoding"},
			{"atbash", "Transformation"},
			{"sha384", "Hashing"},
			{"sha512", "Hashing"},
			{"camel", "Transformation"},
			{"crc32", "Hashing"},
			{"kebab", "Transformation"},
			{"lower", "Transformation"},
			{"upper", "Transformation"},
			{"morse", "Encoding"},
			{"punycode", "Encoding"},
			{"reverse", "Transformation"},
			{"rot13", "Cyphering"},
			{"rot18", "Cyphering"},
			{"rot47", "Cyphering"},
			{"unicode", "Encoding"},
			{"md2", "Hashing"},
			{"sha1", "Hashing"},
		},
	)
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
		return "", fmt.Errorf("invalid index range")
	}

	return input[:start] + encode(input[start:end], method) + input[end:], nil
}

func encode(input, method string) string {
	switch strings.ToLower(method) {
	case "html":
		return jjHtml.Encode(input)
	case "ascii":
		return jjAscii.Encode(input)
	case "hex":
		return jjHex.Encode(input)
	case "binary":
		return jjBinary.Encode(input)
	case "url":
		return jjUrl.Encode(input)
	case "octal":
		return jjOctal.Encode(input)
	case "base32":
		return jjB32.Encode(input)
	case "base45":
		return jjB45.Encode(input)
	case "base64":
		return jjB64.Encode(input)
	case "gzip":
		return jjGzip.Encode(input)
	case "md5":
		return jjMd5.Encode(input)
	case "sha256":
		return jjSha256.Encode(input)
	case "ascii85":
		return jjAscii85.Encode(input)
	case "atbash":
		return jjAtbash.Encode(input)
	case "sha384":
		return jjSha384.Encode(input)
	case "sha512":
		return jjSha512.Encode(input)
	case "camel":
		return jjCamel.Encode(input)
	case "crc32":
		return jjCrc32.Encode(input)
	case "kebab":
		return jjKebab.Encode(input)
	case "lower":
		return jjLower.Encode(input)
	case "upper":
		return jjUpper.Encode(input)
	case "morse":
		return jjMorse.Encode(input)
	case "punycode":
		return jjPunycode.Encode(input)
	case "reverse":
		return jjReverse.Encode(input)
	case "rot13":
		return jjRot13.Encode(input)
	case "rot18":
		return jjRot18.Encode(input)
	case "rot47":
		return jjRot47.Encode(input)
	case "unicode":
		return jjUnicode.Encode(input)
	case "md2":
		return jjMd2.Encode(input)
	case "sha1":
		return jjSha1.Encode(input)

	default:
		showMethodsList()
		return ""
	}
}
