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

	//fs := http.FileServer(http.Dir("static/images"))
	//http.Handle("/Images/", http.StripPrefix("/Images/", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/movies", moviesPage)
	http.HandleFunc("/signin", signinPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
