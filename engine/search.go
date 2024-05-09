package engine

import (
	"log"
	"time"
)

func SearchEngine(query string, idx Index, docs []Document) {
	start := time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("docID: %d\n docText:%s\n", id, doc.Text)
	}
}
