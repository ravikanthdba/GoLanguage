package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/wiki", wiki)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.Header().Set("Status", "HTTP 1/1 200 OK")
	io.WriteString(w, `<img src="/wiki" >`)
}

func wiki(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("Wikipedia-logo-en-big.png")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer file.Close()


	fi, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	http.ServeContent(w, r, file.Name(), fi.ModTime(), file)
}
