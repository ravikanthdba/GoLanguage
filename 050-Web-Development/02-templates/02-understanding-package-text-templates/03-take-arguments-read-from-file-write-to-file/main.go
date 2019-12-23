package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	file, err := template.ParseFiles("input.txt")
	if err != nil {
		log.Fatal("Error reading the file, exiting..")
	}

	newFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal("There is an error in creating the file, exiting..")
	}
	defer newFile.Close()

	if len(os.Args) != 2 {
		switch {
		case len(os.Args) > 2:
			log.Fatal("Excess arguments are passed..")
		case len(os.Args) < 2:
			log.Fatal("Not enough arguments..")
		}
	}
	fmt.Printf("The arguments are %v and %v\n", os.Args[0], os.Args[1])
	name := os.Args[1]

	err = file.Execute(newFile, name)
	if err != nil {
		log.Fatal("Unable to write to the file")
	}
}
