package main

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func dbConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "golang:gotutorial@tcp(127.0.0.1:3306)/tutorial2")
	if err != nil {
		log.Fatalln("Database not connected")
	}
	fmt.Println("DB connected")
	return db
}

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", remove)
	http.HandleFunc("/drop", drop)

	http.ListenAndServe(":8080", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	defer db.Close()
	query, err := db.Prepare(`create table namesList (names varchar(20))`)
	if err != nil {
		log.Fatalln("Unable to Execute the query: ", err)
	}
	result, err := query.Exec()
	if err != nil {
		log.Fatalln(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "Table Created, and Rows Impacted are: %v\n", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	defer db.Close()

	query, err := db.Prepare(`INSERT INTO namesList values ("Ravi")`)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := query.Exec()
	if err != nil {
		log.Fatalln(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "Inserted Ravi into the DB. The rows impacted are: %d\n", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	defer db.Close()
	result, err := db.Query(`select * from namesList`)
	if err != nil {
		log.Fatalln(err)
	}
	defer result.Close()
	fmt.Println("Select ran fine")

	var s, names string
	s += `Record List is: ` + "\n"
	for result.Next() {
		err := result.Scan(&names)
		if err != nil {
			log.Fatalln(err)
		}
		s += names + "\n"
	}
	fmt.Fprintf(w, s)
}

func update(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	defer db.Close()
	query, err := db.Prepare(`update namesList set names = "Raj" where names = "Ravi"`)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := query.Exec()
	if err != nil {
		log.Fatalln(err)
	}

	n, _ := result.RowsAffected()
	fmt.Fprintf(w, "The name has been chaned from Ravi to Raj, the records udpated are: %d\n", n)
}

func remove(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	defer db.Close()
	query, err := db.Prepare(`delete from namesList where names="Raj"`)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := query.Exec()
	if err != nil {
		log.Fatalln(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "Deleted the record with name Raj. The number of records impacted are: %d\n", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	defer db.Close()
	query, err := db.Prepare(`drop table namesList`)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := query.Exec()
	if err != nil {
		log.Fatalln(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "Table dropped, Rows impacted are: %d\n", n)
}


