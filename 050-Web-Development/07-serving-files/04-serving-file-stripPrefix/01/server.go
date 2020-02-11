package main

import (
	"io"
	"net/http"
)

func main() {
	 http.HandleFunc("/wiki", wiki)
	 http.Handle("/images", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))

	 http.ListenAndServe(":8080", nil)
}

func wiki(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Key", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/images/Wikipedia-logo-en-big.png"`)
}


