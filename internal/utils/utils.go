package utils

import (
	"strings"
)

const Base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
const Base = 64

func ReverseString(code string) (reverseCode string) {

	for _, v := range code {
		reverseCode = string(v) + reverseCode
	}

	return reverseCode
}

func Base64Value(char byte) int {
	return strings.IndexByte(string(Base64Chars), char)
}
