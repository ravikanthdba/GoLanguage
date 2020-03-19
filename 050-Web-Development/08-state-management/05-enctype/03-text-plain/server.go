package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	value := make([]byte, r.ContentLength)
	r.Body.Read(value)

	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	var content string = string(value)

	io.WriteString(w, `<form method="post" enctype="text/plain">
		<input type="text" name="firstName" >
		<input type="text" name="lastName" >
		<input type="submit" >
	` + content)
}
