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
}

func main() {
	docs, idx := engine.Initialize()

	r := http.NewServeMux()

	// Create a file server to serve static files
	fs := http.FileServer(http.Dir("static"))
	// Serve static files from the /static/ route
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle requests to the root path by serving the index.html file
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Parse the HTML template
	tmpl := template.Must(template.ParseFiles("static/templates/index.html"))

	// Define a handler function
	r.HandleFunc("POST /templ", func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")
		searchResults := engine.Search(search, idx, docs)

		var data = Data{
			Search:  search,
			Results: searchResults,
		}

		// Execute the template, passing the data
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Println("Listening on port 8000...")
	http.ListenAndServe("localhost:8000", r)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
