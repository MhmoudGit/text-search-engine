package main

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type Document struct {
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocuments(path string) ([]Document, error) {
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

	decode := xml.NewDecoder(gz)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	err = decode.Decode(&dump)
	if err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
