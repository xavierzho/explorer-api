package utils

import (
	"strconv"
)

// M is a type shorthand for param input
type M map[string]string

func Hex2Int(hex string) int64 {
	if hex[:2] == "0x" {
		hex = hex[2:]
	}
	i, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0
	}
	return i
}
