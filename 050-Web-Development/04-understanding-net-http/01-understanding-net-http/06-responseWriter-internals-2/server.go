package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type student struct {
	fname string
	lname string
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

type data struct {
	Method string
	Url *url.URL
	RequestHeader http.Header
	ContentLength int64
	ResponseHeader http.Header
	Submissions url.Values
}

func (s student) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln("Unable to parse the form, failing...", err)
	}

	w.Header().Set("StatusCode", "200")
	w.Header().Set("Status", "OK")
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("WebApplication Response", "This is the WebApplication Response")

	var d = data{
		Method:         r.Method,
		Url:            r.URL,
		RequestHeader:  r.Header,
		ContentLength:  r.ContentLength,
		Submissions: 	r.Form,
		ResponseHeader: w.Header(),
	}


	fmt.Printf("%T\n", d.ResponseHeader)
	fmt.Println(d.ResponseHeader)

	err = t.ExecuteTemplate(w, "index.gohtml", d)
	if err != nil {
		log.Fatalln("Unable to execute the template", err)
	}
}

func main() {
	var s student
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatalln("Unable to listen and serve", err)
	}
}
