package main

import (
	"log"
	"net/http"
	"html/template"
	"net/url"
)

type student struct {
	fname string
	lname string
}

type data struct {
	Method string
	Submissions url.Values
	Url *url.URL
	Header http.Header
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("foo.gohtml"))
}

func (s student) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln("Unable to parse the form", err)
	}

	var d = data{
		Method:      r.Method,
		Submissions: r.Form,
		Url: r.URL,
		Header: r.Header,
	}

	err = t.ExecuteTemplate(w, "foo.gohtml", d)
	if err != nil {
		log.Fatalln("Unable to Execute the template", err)
	}
}

func main() {
	var s student
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatalln("Unable to serve..", err)
	}
}
