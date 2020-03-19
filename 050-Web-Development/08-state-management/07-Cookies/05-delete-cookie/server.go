package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<a href="/set"> Set a Cookie </a>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Session",
		Value: "This is a value",
	})
	fmt.Fprintf(w, `<p> Setting a Cookie, check DEV Tools <a href="/read"> Read the Cookie </a> </p>`)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<p> The value of cookie is: %v </p> <a href="/expire"> Expire </a> `, c)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Session")
	if err != nil {
		fmt.Fprintf(w, `<p> We couldn't fetch the cookie. Click <a href="/set">here</a> to set the cookie </p>`)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
