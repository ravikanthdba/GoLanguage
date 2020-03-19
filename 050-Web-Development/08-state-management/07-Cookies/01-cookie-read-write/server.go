package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", readCookie)

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "My-Cookie",
		Value: "This is a value",
	})
	io.WriteString(w, "Setting a Cookie, check DEV Tools")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("My-Cookie")
	if err != nil {
		http.Error(w, "Error retrieving cookie", 404)
		return
	}
	fmt.Fprintln(w, "Cookie: ", cookie)
}
