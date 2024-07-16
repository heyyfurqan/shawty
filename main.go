package main

import (
	"fmt"
	"furqan/shawty/internal/shortener"
	"furqan/shawty/internal/storage"
)

func main() {
	store := storage.CreateNewStorage()
	var url string

	fmt.Println("Enter a url to shorten: ")
	fmt.Scan(&url)

	indexKey := store.AddUrl(url)

	shortPath := shortener.Encode(indexKey)
	shawtyUrl := "https://shaw.ty/" + shortPath

	fmt.Println("Converted short url is: ", shawtyUrl)

	var shawtyInput string
	fmt.Println("Enter a url to decode: ")
	fmt.Scan(&shawtyInput)
	decodedKey := shortener.Decode(shawtyInput)
	longUrl, ok := store.GetUrl(decodedKey)
	if !ok {
		fmt.Println("Requested url not found :(")
		store.PrintStorage()
	} else {
		fmt.Println("Your decoded url is: ", longUrl)
	}
}
