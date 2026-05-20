package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"

	"codeberg.org/Matthew-K310/matthew-kennedy/pages"
)

//go:embed static/css/* static/assets/* static/badges/*
var staticFiles embed.FS

func main() {
	http.Handle("/", templ.Handler(pages.Home()))
	http.Handle("/now", templ.Handler(pages.Now()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/photos", templ.Handler(pages.Photos()))
	http.Handle("/music", templ.Handler(pages.Music()))
	http.Handle("/contact", templ.Handler(pages.Contact()))
	http.Handle("/links", templ.Handler(pages.Links()))

	staticFS, _ := fs.Sub(staticFiles, "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// handle content files (micro.org, music, other)
	// contentFS, _ := fs.Sub(staticFiles, "content")
	// http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.FS(contentFS))))

	fmt.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
