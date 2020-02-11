package main

import (
	"html/template"
	"log"
	"os"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("input.gohtml"))
}

func main() {
	err := t.ExecuteTemplate(os.Stdout, "input.gohtml", "Ravikanth")
	if err != nil {
		log.Fatal("Error in executing the template")
	}
}
