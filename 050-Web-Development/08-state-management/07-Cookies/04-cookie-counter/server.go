package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/get", getCookie)
	http.ListenAndServe(":8080", nil)
}

var counter int

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "hello-world",
		Value: "This is my value",
	})
	fmt.Fprintf(w, "Cookie has been set")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("hello-world")
	if err != nil {
		log.Fatalln("Error in fetching cookie: ", err)
	} else {
		counter++
	}
	fmt.Fprintf(w, "The counter value is: %v", counter)
}
