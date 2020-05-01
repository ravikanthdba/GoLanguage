package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/index", index)
	http.Handle("/favico.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}



func index(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	s := "Hello World"
	fmt.Fprintf(w, s)
}