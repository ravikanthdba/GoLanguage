package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("foo.gohtml"))
}

func main() {
	var groceriesList = []string{"Grapes", "Banana", "Apple", "PineApple", "Papaya"}

	newfile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Unable to create the file")
	}

	err = t.ExecuteTemplate(newfile, "foo.gohtml", groceriesList)
	if err != nil {
		log.Fatalln("Unable to pass data to template")
	}
}
