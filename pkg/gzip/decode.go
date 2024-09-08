package jjGzip

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

func Decode(input string) string {
	decoded, _ := base64.StdEncoding.DecodeString(input)
	reader, _ := gzip.NewReader(bytes.NewReader(decoded))
	decodedBytes, _ := io.ReadAll(reader)

	return string(decodedBytes)
}
