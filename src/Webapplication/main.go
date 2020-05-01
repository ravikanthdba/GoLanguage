package main

import (
	"encoding/json"
	"fmt"
	"Webapplication/models"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/getUser", getUser)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", nil)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{
		Firstname: "Ravikanth",
		Lastname: "Garimella",
		Age: 32,
		Sex: "Male",
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n",jsonData)
}
