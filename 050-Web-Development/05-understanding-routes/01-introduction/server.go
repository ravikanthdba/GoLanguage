package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/cat": io.WriteString(w, "cat cat cat")
	case "/dog": io.WriteString(w, "dog dog dog")
	case "/": io.WriteString(w, "This is the index page")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
