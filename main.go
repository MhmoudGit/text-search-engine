package main

import (
	"log"
	"time"
)

const dumpPath = "enwiki-latest-abstract1.xml.gz"

func main() {
	log.Println("Full text search in progress...")
	start := time.Now()
	docs, _ := LoadDocuments(dumpPath)
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	SearchEngine("wild cat", idx, docs)
}

func SearchEngine(query string, idx Index, docs []Document) {
	start := time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("docID: %d\n docText:%s\n", id, doc.Text)
	}
}
