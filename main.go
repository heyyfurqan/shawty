package main

import (
	"fmt"
)

func main() {
	indexKey := 69

	shortPath := Encode(indexKey)
	shawtyUrl := "https://shaw.ty/" + shortPath

	fmt.Println("Converted Short Path is: ", shawtyUrl)
}
