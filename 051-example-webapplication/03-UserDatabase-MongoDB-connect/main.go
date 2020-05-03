package main

import (
	"log"
	"net/http"

	"github.com/ravikanthdba/GoLanguage/051-example-webapplication/03-UserDatabase-MongoDB-connect/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	uc := controllers.NewController(getMongoSession())
	http.HandleFunc("/createUser", uc.CreateUser)
	http.HandleFunc("/retrieveUser", uc.RetrieveUser)
	http.HandleFunc("/retrieveUserById", uc.RetrieveUserById)
	http.HandleFunc("/updateUser", uc.UpdateUser)
	http.HandleFunc("/deleteUser", uc.DeleteUser)

	http.ListenAndServe(":8080", nil)
}

func getMongoSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Println(err)
	}

	return s
}
