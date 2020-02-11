package main

import (
	"net/http"
	"html/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index.gohtml", nil)
}

func me(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "me.gohtml", "Ravikanth")
}

func animal(w http.ResponseWriter, r *http.Request) {
	var animalsList = []string{"cat", "dog", "mouse", "goat"}
	t.ExecuteTemplate(w, "animal.gohtml", animalsList)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/animal/", animal)


	http.ListenAndServe(":8080", nil)
}
