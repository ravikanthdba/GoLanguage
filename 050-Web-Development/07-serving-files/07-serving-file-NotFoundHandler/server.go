package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/wiki", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	io.WriteString(w, "Look at terminal")
}
