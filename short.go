package main

import (
	"strings"
)

const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func Encode(key int) string {
	var shortStr strings.Builder

	for key > 0 {
		remainder := key % 64
		shortStr.WriteByte(base64Chars[remainder])
		key = key / 64
	}

	shortCode := reverse(shortStr.String())

	return shortCode
}

func reverse(code string) (reverseCode string) {

	for _, v := range code {
		reverseCode = string(v) + reverseCode
	}

	return reverseCode
}
