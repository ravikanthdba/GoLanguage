package main

import (
	"html/template"
	"net/http"
	"path"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	http.HandleFunc("/", showBooks)
	http.ListenAndServe(":8080", nil)
}

func showBooks(rw http.ResponseWriter, r *http.Request) {
	book := Book{"Building Webapps in GoLang", "Jeremy Saenz"}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(rw, book); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
