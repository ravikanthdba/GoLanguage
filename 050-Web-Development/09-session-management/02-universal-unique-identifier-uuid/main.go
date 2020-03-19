package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"io"
	"log"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.HandleFunc("/deleteCookie", deleteCookie)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the index page")
	_, err := r.Cookie("session")
	if err != nil {
		io.WriteString(w, `<p> This is an Index Page </p>
		<p> Please click <a href="/setCookie"> here </a> to set the cookie </p>`)
		return
	}
		io.WriteString(w, `<p> This is an Index Page </p>
		<p> Please click <a href="/getCookie"> here </a> to get the cookie </p>`)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Setting the cookie")
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalln("Error in Setting the cookie: ", err)
	}
	cookie := &http.Cookie {
		Name:     "session",
		Value:    id.String(),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	io.WriteString(w, `<p> Cookie has been set now </p>
		<p> Please click <a href="/getCookie"> here </a> to get the cookie </p>`)
}


func getCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching the cookie")
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatalln("Cookie has not been set: ", err)
	}
	io.WriteString(w, `<p> The session Cookie is:` + string(cookie.Value) + `</p>
			<p> Please click <a href="/deleteCookie"> here </a> to delete the cookie </p>
			<p> Please click <a href="/index"> here </a> to go back </p>`)
}

func deleteCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting the cookie")
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatalln("Cookie has not been set: ", err)
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	io.WriteString(w, `<p> Cookie has been deleted, please click <a href="/setCookie"> here </a> to set the cookie </p>
					   <p> Please click <a href="/index"> here </a> to go back to Index Page`)
}
