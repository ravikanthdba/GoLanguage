package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("q")
	io.WriteString(w, "Searched string is: " + value)

	anotherValue := r.FormValue("name")
	io.WriteString(w, "Second Searched string is: " + anotherValue)
}
