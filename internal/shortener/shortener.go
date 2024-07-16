package shortener

import (
	"fmt"
	"furqan/shawty/internal/utils"
	"math"
	"net/url"
	"strings"
)

func Encode(key int) string {
	var shortStr strings.Builder

	for key > 0 {
		remainder := key % int(utils.Base)
		shortStr.WriteByte(utils.Base64Chars[remainder])
		key = key / int(utils.Base)
	}

	shortCode := utils.ReverseString(shortStr.String())

	return shortCode
}

func Decode(shortUrl string) int {
	u, err := url.Parse(shortUrl)
	if err != nil {
		fmt.Println(err)
	}

	shortCodeUrl := u.Path
	shortCode, _ := strings.CutPrefix(shortCodeUrl, "/")

	var result float64 = 0

	for j := 0; j < len(shortCode); j++ {
		char := shortCode[len(shortCode)-1-j]
		value := float64(utils.Base64Value(char))
		power := math.Pow(float64(utils.Base), float64(j))
		result += value * power
	}
	return int(result)
}
