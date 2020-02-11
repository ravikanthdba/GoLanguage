package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/wiki", wiki)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Header().Set("Status", "HTTP 1/1 200 OK")
	io.WriteString(w, `<img src="wiki">`)
}

func wiki(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Wikipedia-logo-en-big.png")
}
