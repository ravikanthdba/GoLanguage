package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"log"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/index", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/home", home)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyprocess)

	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "home.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyprocess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Fatalln("The error for the request: ", w, " is: ", err)
	}
}
