package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	name string
	age  int
}

func main() {
	file, err := template.ParseFiles("input.gohtml")
	if err != nil {
		log.Fatal("Unable to parse the file: input.gohtml")
	}

	err = file.Execute(os.Stdout, "Ravikanth")
	if err != nil {
		log.Fatalln("Unable to execute, hence failing")
	}
}
