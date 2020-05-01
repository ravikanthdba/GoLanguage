package main

import (
	"UserDatabase-CodeOrganized/controllers"
	"net/http"
)

func main() {
	uc := controllers.NewController();
	http.HandleFunc("/createUser", uc.CreateUser)
	http.HandleFunc("/retrieveUser", uc.RetrieveUser)
	http.HandleFunc("/retrieveUserById", uc.RetrieveUserById)
	http.HandleFunc("/updateUser", uc.UpdateUser)
	http.HandleFunc("/deleteUser", uc.DeleteUser)

	http.ListenAndServe(":8080", nil)
}
