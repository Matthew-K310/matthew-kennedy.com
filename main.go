package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := Home()

	http.Handle("/", templ.Handler(component))

	// Serve files from ./static at /static/
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)

}
