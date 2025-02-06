package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/Webflix.html")
}

func moviesPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func signinPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/signin.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/Contact Us.html")
}

func main() {

	http.HandleFunc("/home", homePage)
	http.HandleFunc("/movies", moviesPage)
	http.HandleFunc("/signin", signinPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
