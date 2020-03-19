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
	body := string(value)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<form method="post" enctype="application/x-www-form-urlencoded">
	<input name="firstName" type="text">
	<input name="lastName" type="text">
	<input type="submit">
	` + body)
}
