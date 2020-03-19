package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method at index is: ", r.Method)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Status", "200 OK")
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method at foo is: ", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
}
