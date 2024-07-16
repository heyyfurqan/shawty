package storage

import "fmt"

type UrlStore struct {
	urlMap     map[int]string
	lastUsedID int
}

func CreateNewStorage() *UrlStore {
	return &UrlStore{
		urlMap:     make(map[int]string),
		lastUsedID: 0,
	}
}

func (s *UrlStore) AddUrl(url string) int {
	s.lastUsedID++
	s.urlMap[s.lastUsedID] = url
	return s.lastUsedID
}

func (s *UrlStore) GetUrl(id int) (string, bool) {
	url, exists := s.urlMap[id]
	return url, exists
}

func (s *UrlStore) PrintStorage() {
	for key, value := range s.urlMap {
		fmt.Printf("%d is the id for the url %q\n", key, value)
	}
}
