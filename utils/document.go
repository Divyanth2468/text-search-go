package utils

import (
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
)

type Entries struct {
	Documents []Document `xml:"entry"`
}

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int    `xml:"-"`
}

type Cache struct {
	Hash      string     `json:"hash"`
	Documents []Document `json:"documents"`
}

func LoadDocuments(path string) ([]Document, error) {

	hash, err := hashFile(path)

	if err != nil {
		return nil, err
	}

	cachePath := filepath.Base(path) + ".cache.json"
	log.Println(cachePath)
	if _, err := os.Stat(cachePath); err == nil {
		cacheFile, err := os.Open(cachePath)
		if err == nil {
			defer cacheFile.Close()
			var cache Cache
			if err := json.NewDecoder(cacheFile).Decode(&cache); err == nil && cache.Hash == hash {
				for i := range cache.Documents {
					cache.Documents[i].ID = i
				}
				log.Println("Used Cached file")
				return cache.Documents, nil
			}
		}
	}

	// Fall back to fresh parse
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}

	defer gz.Close()

	dec := xml.NewDecoder(gz)

	var entries Entries

	if err := dec.Decode(&entries); err != nil {
		return nil, err
	}

	for i := range entries.Documents {
		entries.Documents[i].ID = i
	}

	cacheData := Cache{
		Hash:      hash,
		Documents: entries.Documents,
	}

	cacheOut, err := os.Create(cachePath)
	if err == nil {
		defer cacheOut.Close()
		_ = json.NewEncoder(cacheOut).Encode(cacheData)
	}

	log.Println("Didnt use cache")
	return entries.Documents, nil
}
