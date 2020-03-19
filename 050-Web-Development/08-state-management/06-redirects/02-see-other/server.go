package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/bar", bar)

	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method at helloWorld is: ", r.Method)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Status", "200 OK")
	io.WriteString(w, "Hello World")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method at foo is: ", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}
