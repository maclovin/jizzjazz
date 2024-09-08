package jjCrc32

import (
	"encoding/hex"
	"hash/crc32"
)

func Encode(input string) string {
	checksum := crc32.ChecksumIEEE([]byte(input))

	return hex.EncodeToString([]byte{byte(checksum >> 24), byte(checksum >> 16), byte(checksum >> 8), byte(checksum)})
}
