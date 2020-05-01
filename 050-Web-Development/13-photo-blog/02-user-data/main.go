package main

import (
	"html/template"
	"log"
	"net/http"
	uuid "github.com/satori/go.uuid"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}


func main() {
	http.HandleFunc("/index", index)
	http.Handle("/favico.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", cookie)
}


func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		log.Println("Cookie is not set, setting the cookie")
		id, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:       "session",
			Value:      id.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
