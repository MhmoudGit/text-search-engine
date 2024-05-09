package engine

import (
	"log"
	"time"
)

func Initialize() ([]Document, Index) {
	const dumpPath = "enwiki-latest-abstract1.xml.gz"
	log.Println("Full text search in progress...")
	start := time.Now()
	docs, _ := LoadDocuments(dumpPath)
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start).Milliseconds())

	return docs, idx
}

func Search(query string, idx Index, docs []Document) []Document {
	var result []Document
	start := time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		// log.Printf("docID: %d\n docText:%s\n", id, doc.Text)
		result = append(result, doc)
	}
	return result
}
