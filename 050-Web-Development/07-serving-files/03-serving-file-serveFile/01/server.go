package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/wiki", wiki)

	http.ListenAndServe(":8080", nil)
}

func wiki(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="Wikipedia-logo-en-big.png" >`)
}
