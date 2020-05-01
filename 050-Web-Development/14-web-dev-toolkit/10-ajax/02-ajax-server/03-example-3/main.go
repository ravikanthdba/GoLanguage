package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/validateUser", validateUser)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", nil)
}

func validateUser(w http.ResponseWriter, r *http.Request) {
	var usersList = map[string]bool{
		"Ravikanth": true,
		"Nagabhushanam": true,
		"Swarna Latha": true,
		"Viraj Nandan": true,
		"Bhargavi": true,
	}

	str, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Unable to read request body", err)
	}

	fmt.Println("Checking the user: ", string(str), " and username: " ,usersList[string(str)])
	fmt.Fprintln(w, usersList[string(str)])
}