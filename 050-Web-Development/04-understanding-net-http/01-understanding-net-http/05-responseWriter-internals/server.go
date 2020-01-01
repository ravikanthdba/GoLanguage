package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Ravikanth-Key", "This is from the Ravikanth webapp")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> This is an example for Response Writer </h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
