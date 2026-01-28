package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"

	"github.com/Matthew-K310/matthew-kennedy/pages"
)

//go:embed static/css/* static/assets/*
var staticFiles embed.FS

func main() {
	http.Handle("/", templ.Handler(pages.Home()))
	http.Handle("/now", templ.Handler(pages.Now()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/micro", templ.Handler(pages.Micro()))
	http.Handle("/photos", templ.Handler(pages.Photos()))
	http.Handle("/music", templ.Handler(pages.Music()))

	// Serve embedded static files at /static/
	staticFS, _ := fs.Sub(staticFiles, "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
