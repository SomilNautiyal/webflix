package main

import (
	"log"
	"net/http"
)

// Serve static files directly from "static/" without needing "/static/" in URLs
func serveStaticFiles() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
}

func main() {
	serveStaticFiles() // Setup static file handling

	// Start the server
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
