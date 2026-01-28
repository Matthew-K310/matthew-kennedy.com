package main

import (
	"fmt"
	"net/http"

	"github.com/Matthew-K310/matthew-kennedy/pages"
	"github.com/a-h/templ"
)

func main() {
	// Assuming you have a Home() function in pages package
	http.Handle("/", templ.Handler(pages.Home()))
	http.Handle("/now", templ.Handler(pages.Now()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/micro", templ.Handler(pages.Micro()))
	http.Handle("/photos", templ.Handler(pages.Photos()))
	http.Handle("/music", templ.Handler(pages.Music()))

	// Serve files from ./static at /static/
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
