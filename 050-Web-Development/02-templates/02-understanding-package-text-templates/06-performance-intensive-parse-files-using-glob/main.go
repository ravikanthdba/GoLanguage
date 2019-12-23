package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	fmt.Println("Executing the initialization function..")
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("Executing the main function")
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
