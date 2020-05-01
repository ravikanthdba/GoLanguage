package main

import (
	"UserDatabase-MongoDB-connect/controllers"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	uc := controllers.NewController(getMongoSession());
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
