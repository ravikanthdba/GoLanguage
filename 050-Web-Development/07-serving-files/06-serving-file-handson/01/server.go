package main

import (
	"html/template"
	"io"
	"net/http"
)
var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("foo.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/wiki", wiki)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func wiki(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "foo.gohtml", "https://i.dlpng.com/static/png/1194040-labrador-retriever-siberian-husky-pug-puppy-cat-puppy-puppy-png-900_800_preview.png")
}


