package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an index page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an about page")
}

func help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an help page")
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/about", http.HandlerFunc(about))
	http.Handle("/help", http.HandlerFunc(help))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
