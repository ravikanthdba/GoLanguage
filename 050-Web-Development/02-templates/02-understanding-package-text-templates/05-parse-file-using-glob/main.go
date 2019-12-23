package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	file, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatal("Error parsing the files")
	}

	err = file.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatal("Error executing the template one.gohtml")
	}

	err = file.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatal("Error executing the template three.gohtml")
	}

	err = file.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal("Error executing the template two.gohtml")
	}

	err = file.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error executing the template")
	}
}
