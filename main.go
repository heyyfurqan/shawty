package main

import (
	"fmt"
)

func main() {
	indexKey := 7465

	shortPath := Encode(indexKey)
	shawtyUrl := "https://shaw.ty/" + shortPath

	fmt.Println("Converted Short Path is: ", shawtyUrl)

	decodedKey := Decode(shortPath)
	fmt.Println("Decoded back is: ", decodedKey)

	if decodedKey == indexKey {
		fmt.Println("It works.")
	} else {
		fmt.Println("Nuh uh.")
	}
}
