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
	content := string(value)

	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	io.WriteString(w, `<form method="post" enctype="multipart/form-data">
		<input type="text" name="firstName">
		<input type="text" name="lastName">
		<input type="submit">
		<br />
    ` + content)
}
