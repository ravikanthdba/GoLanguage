package main

import (
	"17-Databases-Code-Organization/controllers"
	"net/http"
)

func main() {
	employee := controllers.NewController()
	http.HandleFunc("/createData", employee.CreateData)
	http.HandleFunc("/retrieveData", employee.RetrieveData)
	http.HandleFunc("/updateData", employee.UpdateData)
	http.HandleFunc("/deleteData", employee.DeleteData)

	http.ListenAndServe(":8080", nil)
}
