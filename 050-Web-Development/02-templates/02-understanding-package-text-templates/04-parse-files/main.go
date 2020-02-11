package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	file, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	file, err = file.ParseFiles("two.gohtml", "three.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = file.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = file.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = file.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
