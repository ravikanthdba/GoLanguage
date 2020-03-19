package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", StatusCode)
	http.HandleFunc("/index", index)

	http.ListenAndServe(":8080", nil)
}


func StatusCode(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GOOD")
}

func index(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post">
			<input type="text" name="q">
			<input type="submit">
		</form>
	` + v)
}

