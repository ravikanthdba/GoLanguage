package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	file, err := template.ParseFiles("main.gohtml")
	if err != nil {
		log.Fatal("File not found, exiting")
	}

	err = file.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Failed to write the data")
	}
}
