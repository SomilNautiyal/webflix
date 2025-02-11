package main

import (
	"net/http"
)

// Capitalize HomePage to make it accessible in tests
func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/Webflix.html")
}

func main() {
	http.HandleFunc("/", HomePage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}
