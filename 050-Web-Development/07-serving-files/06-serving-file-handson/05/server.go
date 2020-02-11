package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/index", index)
	http.HandleFunc("/apply", apply)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "home.gohtml", "Home Method")
}

func about(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "about.gohtml", "About Method")
}

func index(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index.gohtml", "Index Method")
}

func apply(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		t.ExecuteTemplate(w, "applyPost.gohtml", "Post Apply Method")
		return
	}
	t.ExecuteTemplate(w, "apply.gohtml", "Apply Method")
}
