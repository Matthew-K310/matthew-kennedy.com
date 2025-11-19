package main

import (
	"github.com/Matthew-K310/matthew-kennedy/internal/templates"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts := []templates.Post{
			{Title: "Hello World", Slug: "hello-world", Date: "2025-01-01"},
		}
		templates.Home(posts).Render(r.Context(), w)
	})

	log.Println("running on :8080")
	http.ListenAndServe(":8080", mux)
}
