package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ravikanthdba/GoLanguage/051-example-webapplication/02-UserDatabase-CodeOrganized/models"
)

type UserController struct{}

var userDB []models.User

func NewController() *UserController {
	return &UserController{}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	jsonData, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Println(err)
	}
	userDB = append(userDB, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonData)
	log.Println("User added to the Database")
}

func (uc UserController) RetrieveUser(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(userDB)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonData)
	log.Println("Data Retrieved")
}

func (uc UserController) RetrieveUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idValue, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}

	for _, data := range userDB {
		if data.Id == idValue {
			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Println(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s\n", jsonData)
			log.Println("Data By ID Retrieved")
		}
	}
	w.WriteHeader(http.StatusNotFound)
	log.Println("Data Not Retrieved")
}

func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var newData []models.User
	id := r.URL.Query().Get("id")
	idValue, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	for _, data := range userDB {
		if data.Id != idValue {
			newData = append(newData, data)
		} else {
			newData = append(newData, user)
		}
	}
	userDB = newData
	jsonData, err := json.MarshalIndent(userDB, "", " ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jsonData)
	log.Println("Updation Completed")
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idValue, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}

	var newDB []models.User
	for _, data := range userDB {
		if data.Id != idValue {
			newDB = append(newDB, data)
		}
	}

	userDB = newDB
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.MarshalIndent(userDB, "", " ")
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "%s\n", jsonData)
	log.Println("Deletion Completed")
}
