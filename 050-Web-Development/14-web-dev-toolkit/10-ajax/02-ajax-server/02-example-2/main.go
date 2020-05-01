package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)

	http.ListenAndServe(":8080", nil)
}

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", nil)
}


func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}