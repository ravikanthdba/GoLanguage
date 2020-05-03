package main

import (
	"net/http"

	"github.com/ravikanthdba/GoLanguage/051-example-webapplication/02-UserDatabase-CodeOrganized/controllers"
)

func main() {
	uc := controllers.NewController()
	http.HandleFunc("/createUser", uc.CreateUser)
	http.HandleFunc("/retrieveUser", uc.RetrieveUser)
	http.HandleFunc("/retrieveUserById", uc.RetrieveUserById)
	http.HandleFunc("/updateUser", uc.UpdateUser)
	http.HandleFunc("/deleteUser", uc.DeleteUser)

	http.ListenAndServe(":8080", nil)
}
