package main

import (
	"html/template"
	"log"
	"net/http"
)

type student struct {
	fname string
	lname string
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("foo.gohtml"))
}

func (s student) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln("Error in parsing the form")
	}

	err = t.ExecuteTemplate(w, "foo.gohtml", r.Form)
}

func main() {
	var s student
	http.ListenAndServe(":8080", s)
}
