package main

import (
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/foo.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("starting-files"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute")
	}
}

