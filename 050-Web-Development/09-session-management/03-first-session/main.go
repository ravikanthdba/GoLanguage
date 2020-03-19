package main

import (
	"fmt"
	"html/template"
	"net/http"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Firstname string
	Lastname string
	Username string
}

var sessionDB = make(map[string]string)
var userDB = make(map[string]User)

var t *template.Template

func init() {
	t = template.Must(t.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	fmt.Println("Cookie from foo is: ", cookie)
	if err != nil {
		fmt.Println("Cookie is not set, setting it now")
	var user User
		if r.Method == http.MethodPost {
			id, _ := uuid.NewV4()
			cookie := &http.Cookie{
				Name:       "session",
				Value:      id.String(),
			}

			http.SetCookie(w, cookie)
			user.Firstname = r.FormValue("FirstName")
			user.Lastname = r.FormValue("LastName")
			user.Username = r.FormValue("Username")

			sessionDB[cookie.Value] = user.Username
			userDB[user.Username] = user

			fmt.Println("userDB is: ", userDB)
			fmt.Println("sessionDB is: ", sessionDB)
		}
		t.ExecuteTemplate(w, "foo.gohtml", user)
	} else {
		fmt.Println("Cookie is set already")
		username := sessionDB[cookie.Value]
		Userdetails := userDB[username]
		fmt.Println(username)
		fmt.Println(Userdetails)
		t.ExecuteTemplate(w, "foo.gohtml", Userdetails)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/foo", http.StatusSeeOther)
		return
	}

	username := sessionDB[cookie.Value]
	Userdetails := userDB[username]

	fmt.Println("From the Bar function, username = ", username, " and userdetails = ", Userdetails)
	t.ExecuteTemplate(w, "bar.gohtml", Userdetails)
}