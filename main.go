package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/Matthew-K310/matthew-kennedy/internal/micro"
	"github.com/Matthew-K310/matthew-kennedy/pages"
)

//go:embed static/css/* static/assets/*
var staticFiles embed.FS

func main() {
	posts, err := micro.LoadPosts("content/micro.org")
	if err != nil {
		log.Fatalf("failed to load micro posts: %v", err)
	}

	http.Handle("/", templ.Handler(pages.Home()))
	http.Handle("/now", templ.Handler(pages.Now()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/micro", templ.Handler(pages.Micro(posts)))
	http.Handle("/photos", templ.Handler(pages.Photos()))
	http.Handle("/music", templ.Handler(pages.Music()))

	staticFS, _ := fs.Sub(staticFiles, "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// handle content files (micro.org)
	contentFS, _ := fs.Sub(staticFiles, "content")
	http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.FS(contentFS))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
