package main

import (
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Ravikanth!!")
}


func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a dog function!!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/", dog)

	http.ListenAndServe(":8080", nil)
}
