package main

import (
	"log"
	"time"

	"github.com/MhmoudGit/text-search-engine/engine"
)

const dumpPath = "enwiki-latest-abstract1.xml.gz"

func main() {
	log.Println("Full text search in progress...")
	start := time.Now()
	docs, _ := engine.LoadDocuments(dumpPath)
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(engine.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	engine.SearchEngine("wild cat", idx, docs)
}
