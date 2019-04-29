package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	// Main endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World, this website is being served by GoLang")
	})
	// Example Endpoint
	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the /example endpoint")
	})

	// Serving the static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":"+PORT, nil)
}
