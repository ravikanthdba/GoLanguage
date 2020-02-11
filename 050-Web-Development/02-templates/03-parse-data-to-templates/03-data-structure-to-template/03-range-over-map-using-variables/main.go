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
	var grocieriesList = map[string][]string{"Vegetables": []string{"Brinjal", "Carrot", "Potato"},
		"Fruits": []string{"Tomato", "Pomegranete", "Apple", "Pineapple"},
	}
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("The output file cannot be created, exiting")
	}

	err = t.ExecuteTemplate(newFile, "index.gohtml", grocieriesList)
	if err != nil {
		log.Fatalln("Cannot pass variables to the template")
	}

}
