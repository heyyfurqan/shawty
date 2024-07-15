package main

import (
	"fmt"
	"math"
	"net/url"
	"strings"
)

const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
const base = 64

func Encode(key int) string {
	var shortStr strings.Builder

	for key > 0 {
		remainder := key % int(base)
		shortStr.WriteByte(base64Chars[remainder])
		key = key / int(base)
	}

	shortCode := reverse(shortStr.String())

	return shortCode
}

func Decode(shortUrl string) int {
	u, err := url.Parse(shortUrl)
	if err != nil {
		fmt.Println(err)
	}

	shortCode := u.Path

	var result float64 = 0

	for j := 0; j < len(shortCode); j++ {
		char := shortCode[len(shortCode)-1-j]
		value := float64(base64Value(char))
		power := math.Pow(float64(base), float64(j))
		result += value * power
	}
	return int(result)
}

func reverse(code string) (reverseCode string) {

	for _, v := range code {
		reverseCode = string(v) + reverseCode
	}

	return reverseCode
}

func base64Value(char byte) int {
	return strings.IndexByte(string(base64Chars), char)
}
