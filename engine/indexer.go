package engine

import (
	"encoding/json"
	"os"
)

func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

type Index map[string][]int

func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		for _, token := range Analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func (idx *Index) GetData() error {
	// Open the JSON file for reading
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON data into the variable
	err = decoder.Decode(&idx)
	if err != nil {
		return err
	}

	return nil
}

func (idx Index) StoreData() error {
	// save the indexed data inside a json file
	// Open a new JSON file for writing
	file, err := os.Create("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	// Encode the data and write it to the file
	err = encoder.Encode(idx)
	if err != nil {
		return err
	}
	return nil
}

func (idx Index) Search(text string) []int {
	var r []int
	for _, token := range Analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}
