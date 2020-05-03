package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/ravikanthdba/GoLanguage/051-user-database-webapplication/01-UserDatabase/models"
)

var t *template.Template
var usersDatabase []models.UserDB

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/retrieveUser", retrieveUser)
	http.HandleFunc("/updateUser", updateUser)
	http.HandleFunc("/deleteUser", deleteUser)
	http.HandleFunc("/listUser", listUser)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := models.UserDB{}
	json.NewDecoder(r.Body).Decode(&user)
	usersDatabase = append(usersDatabase, user)
	jsonData, _ := json.MarshalIndent(user, "", " ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", jsonData)
}

func retrieveUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, data := range usersDatabase {
		idValue, _ := strconv.Atoi(id)
		if data.Id == idValue {
			jsondata, err := json.MarshalIndent(data, "", " ")
			if err != nil {
				log.Println(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s\n", jsondata)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%s\n", "No user ID exists")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	user := models.UserDB{}
	json.NewDecoder(r.Body).Decode(&user)
	id := r.URL.Query().Get("id")
	idValue, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}

	var newDatabase []models.UserDB

	for _, data := range usersDatabase {
		if data.Id != idValue {
			newDatabase = append(newDatabase, data)
		} else if data.Id == idValue {
			newDatabase = append(newDatabase, user)
		}
	}

	usersDatabase = newDatabase
	jsonData, err := json.MarshalIndent(usersDatabase, "", " ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonData)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idValue, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}

	var newDatabase []models.UserDB

	for _, data := range usersDatabase {
		if data.Id != idValue {
			newDatabase = append(newDatabase, data)
		}
	}
	usersDatabase = newDatabase
	jsonData, err := json.MarshalIndent(usersDatabase, "", " ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonData)
}

func listUser(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.MarshalIndent(usersDatabase, "", " ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonData)
}
