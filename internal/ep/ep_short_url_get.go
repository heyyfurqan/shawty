package ep

import (
	"encoding/json"
	shortener "furqan/shawty/internal/shortener"
	storage "furqan/shawty/internal/storage"
	utils "furqan/shawty/internal/utils"
	"net/http"
)

type ShortenUrlReq struct {
	Url string `json:"url"`
}

type ShortenUrlResp struct {
	ShortUrl string `json:"shortUrl"`
}

func ShortenUrl(w http.ResponseWriter, r *http.Request, db *storage.UrlStore) {
	var body ShortenUrlReq

	// Get url from body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get index key from storage
	indexKey := db.AddUrl(body.Url)

	// Encode index key.
	shortCode := shortener.Encode(indexKey)

	shortUrl := utils.ShawtyUrl + shortCode

	json.NewEncoder(w).Encode(shortUrl)
}
