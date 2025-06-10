package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	utils "github.com/Divyanth2468/full-text-search-engine/utils"
)

func main() {

	// Loading file

	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "/Users/uppuluridivyanthsatya/Desktop/Go/full-text-search-engine/dbpedia_abstracts.xml.gz", "Abstract Dump path")
	flag.StringVar(&query, "q", "a car on road going to hyderabad with a man driving it", "search query")
	flag.Parse()
	log.Println("Full text search is in progress")
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
	start = time.Now()

	// Indexing
	if _, err := os.Stat("Cache"); err != nil {
		if err := os.MkdirAll("Cache", 0755); err != nil {
			panic("Not able to create Cache directory")
		}
	}
	indexPath := filepath.Join("Cache", filepath.Base(dumpPath)+".index_cache.json")
	var idx utils.Index = make(utils.Index)
	if _, err := os.Stat(indexPath); err == nil {
		idx, err = utils.LoadIndexFromFile(indexPath)
		if err != nil {
			panic("Failed to load index " + err.Error())
		}
		log.Println("Loaded from cached index")
	} else {
		idx.Add(docs)
		err := idx.SaveToFile(indexPath)
		if err != nil {
			panic("Failed to save index: " + err.Error())
		}
	}
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	// Search

	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
