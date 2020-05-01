package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type person struct {
	FirstName string
	Lastname string
	Age int
}


func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/marshal", marshalling)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You are at Index page")
}

func marshalling(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		FirstName: "Ravikanth",
		Lastname:  "Garimella",
		Age:       32,
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("status", "200 OK")
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
