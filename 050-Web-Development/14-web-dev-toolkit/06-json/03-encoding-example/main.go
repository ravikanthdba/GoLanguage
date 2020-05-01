package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type person struct {
	Firstname string
	Lastname string
	Age string
	Hobbies []string
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/encoding", encodingJson)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You are at index page")
}

func encodingJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Status", "200 OK")

	p1 := &person{
		Firstname: "Ravikanth",
		Lastname:  "Garimella",
		Age:       "32",
		Hobbies: []string{"Playing cricket", "Coding"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
