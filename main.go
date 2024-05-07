package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	var dumpPath, query string

	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract data dump")
	flag.StringVar(&query, "q", "wild cat", "search query")
	flag.Parse()
	log.Println("Full text search in progress...")

	start := time.Now()
	docs, _ := LoadDocuments(dumpPath)
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("docID: %d\n docText:%s\n", id, doc.Text)
	}
}
