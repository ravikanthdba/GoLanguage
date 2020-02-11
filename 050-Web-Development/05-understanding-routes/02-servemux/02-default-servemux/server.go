package main

import (
	"io"
	"net/http"
)

type cat int
type dog int

func (c cat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat cat cat")
}

func (d dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog dog dog")
}

func main() {
	var c cat
	var d dog

	http.Handle("/cat", c)
	http.Handle("/dog/", d)

	http.ListenAndServe(":8080", nil)
}
