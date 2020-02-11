package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var groceriesList = []string{"Grapes", "Banana", "Apple", "PineApple", "Papaya"}

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating the file, exiting..")
	}

	err = t.ExecuteTemplate(newFile, "index.gohtml", groceriesList)
	if err != nil {
		log.Fatalln("Error passing the data to the template")
	}
}
