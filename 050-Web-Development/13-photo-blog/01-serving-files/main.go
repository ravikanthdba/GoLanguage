package main

import (
	"html/template"
	"net/http"
	uuid "github.com/satori/go.uuid"
)

var t *template.Template

func init() {
	t = template.Must(t.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/index", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", cookie)
}


func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		id,_ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:       "session",
			Value:      id.String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}
