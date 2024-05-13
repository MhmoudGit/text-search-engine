package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MhmoudGit/text-search-engine/engine"
)

type Data struct {
	Search  string
	Results []engine.Document
	Len     int
}

func main() {
	docs, idx := engine.Initialize()

	mux := http.NewServeMux()

	// Create a file server to serve static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	// Parse the HTML template
	tmpl := template.Must(template.ParseFiles("static/templates/index.html"))

	// Handlers
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	mux.HandleFunc("POST /search", func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")
		searchResults := engine.Search(search, idx, docs)

		var data = Data{
			Search:  search,
			Results: searchResults,
			Len: len(searchResults),
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Println("Listening on port 8000")
	http.ListenAndServe("localhost:8000", mux)
}
