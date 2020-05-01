package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,  "Hello Golang from Container")
}
