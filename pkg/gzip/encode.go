package jjGzip

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
)

func Encode(input string) string {
	var buf bytes.Buffer

	writer := gzip.NewWriter(&buf)
	writer.Write([]byte(input))
	writer.Close()

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
