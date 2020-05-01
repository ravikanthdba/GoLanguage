package main

import (
	"html/template"
	"log"
	"net/http"
	uuid "github.com/satori/go.uuid"
	"strings"
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
	items := getListOfItems(cookie)

	itemsList := strings.Split(items, "|")
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", itemsList)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println("Cookie not set, setting the cookie")
		id,_ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:       "session",
			Value:      id.String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func getListOfItems(c *http.Cookie) string {
	s := c.Value + "|"
	item_1 := "Brinjal"
	item_2 := "Onion"
	item_3 := "Tomato"

	if !strings.Contains(s, item_1) {
		s += item_1 + "|"
	}

	if !strings.Contains(s, item_2) {
		s += item_2 + "|"
	}

	if !strings.Contains(s, item_3) {
		s += item_3 + "|"
	}

	return s
}


