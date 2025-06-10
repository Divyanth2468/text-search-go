package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	dumpPath = strings.TrimSpace(dumpPath)
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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your search queries below (type 'exit' to quit):")

	for {
		if !scanner.Scan() {
			break
		}

		query := strings.TrimSpace(scanner.Text())

		if query == "exit" || query == "quit" {
			break
		}

		if query == "" {
			continue
		}

		// Search

		start = time.Now()
		matchedIDs := idx.Search(query)
		fmt.Printf("Found %d results in %v \n", len(matchedIDs), time.Since(start))

		for _, id := range matchedIDs {
			doc := docs[id]
			fmt.Printf("Title: %s\nURL: %s\nDescription:%s\n--\n", doc.Title, doc.URL, doc.Text)
		}

	}
}
