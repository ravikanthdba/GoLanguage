package controllers

import (
	"17-Databases-Code-Organization/models"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type EmployeeController struct {}
var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("../templates/*"))
}

func dbConnection() (*sql.DB) {
	db, err := sql.Open("mysql", "gouser:databasepassword@/employee")
	if err != nil {
		log.Println("Unable to connect", err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	return db
}

func NewController() *EmployeeController {
	return &EmployeeController{}
}

func (employee EmployeeController) CreateData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		age := r.FormValue("age")
		posting := r.FormValue("posting")
		salary := r.FormValue("salary")
		sex := r.FormValue("gender")

		db := dbConnection()
		defer db.Close()

		insertData, err := db.Prepare("insert into employee (empname, age, posting, salary, sex) values (?, ?, ?, ?, ?)")
		if err != nil {
			log.Println(err)
		}
		insertData.Exec(name, age, posting, salary, sex)
		log.Println("Data for empname: ", name, " created.")
		http.Redirect(w, r, "/retrieveData", http.StatusSeeOther)
	} else {
		t.ExecuteTemplate(w, "createemployee.gohtml", nil)
	}
}


func (employee EmployeeController) RetrieveData(w http.ResponseWriter, r *http.Request) {
	e := models.Employee{}
	var result []models.Employee

	db := dbConnection()
	defer db.Close()
	selectData, err := db.Query("select empid, empname, age, posting, salary, sex from employee")
	if err != nil {
		log.Println(err)
	}

	for selectData.Next() {
		var name, posting, sex string
		var empid, age,salary int
		err := selectData.Scan(&empid, &name, &age, &posting, &salary, &sex)
		if err != nil {
			log.Println(err)
		}

		e.Empid = empid
		e.Empname = name
		e.Age = age
		e.Posting = posting
		e.Salary = salary
		e.Sex = sex

		result = append(result, e)
	}

	log.Println(result)
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "retrieveemployee.gohtml", result)
}


func (employee EmployeeController) UpdateData(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	defer db.Close()
	if r.Method == http.MethodPost {
		key := r.URL.Query().Get("id")
		empname := r.FormValue("empname")
		age := r.FormValue("age")
		posting := r.FormValue("posting")
		salary := r.FormValue("salary")
		sex := r.FormValue("gender")

		updateData, err := db.Prepare("update employee set empname = ?, age = ?, posting = ?, salary = ?, sex = ? where empid = ?")
		if err != nil {
			log.Println(err)
		}
		_, err = updateData.Exec(empname, age, posting, salary, sex, key)
		if err != nil {
			log.Println(err)
		}
		log.Println("Record updated for key: ", key)
		http.Redirect(w, r, "/retrieveData", http.StatusSeeOther)
	} else {
		key := r.URL.Query().Get("id")
		var result []models.Employee
		e := models.Employee{}
		selectData, err := db.Query("select empid, empname, age, posting, salary, sex from employee where empid = " + key)
		if err != nil {
			log.Println(err)
		}

		for selectData.Next() {
			var empid, age, salary int
			var empname, posting, sex string

			err := selectData.Scan(&empid, &empname, &age, &posting, &salary, &sex)
			if err != nil {
				log.Println(err)
			}

			e.Empid = empid
			e.Empname = empname
			e.Age = age
			e.Posting = posting
			e.Salary = salary
			e.Sex = sex


			result = append(result, e)
		}

		t.ExecuteTemplate(w, "updateemployee.gohtml", result)
	}
}


func (employee EmployeeController) DeleteData(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("id")
	db := dbConnection()
	defer db.Close()
	deleteData, err := db.Prepare("delete from employee where empid = ?")
	if err != nil {
		log.Println(err)
	}
	_, err = deleteData.Exec(key)
	log.Println("data with empid: ", key, " has been deleted")
	http.Redirect(w, r, "/retrieveData", http.StatusSeeOther)
}