package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", bar)
	http.HandleFunc("/foo", foo)

	http.ListenAndServe(":8080", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method at foo is: ", r.Method)
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Status", "200 OK")
	io.WriteString(w, "Temporary redirected to bar page")
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method at foo is: ", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
