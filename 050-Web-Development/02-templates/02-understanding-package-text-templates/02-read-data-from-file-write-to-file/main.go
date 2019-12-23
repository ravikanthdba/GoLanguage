package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	file, err := template.ParseFiles("input.gohtml")
	if err != nil {
		log.Fatal("Failed to read data from the file")
	}

	newfile, err := os.Create("input.txt")
	if err != nil {
		log.Fatal("Failed to create the new file, exiting..")
	}
	defer newfile.Close()

	err = file.Execute(newfile, nil)
	if err != nil {
		log.Fatal("Failed to write data to the file", err)
	}

}
