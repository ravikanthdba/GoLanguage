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

type groceries struct {
	Item  string
	Price float64
}

func main() {
	var items []groceries
	g1 := groceries{
		Item:  "onion",
		Price: 100,
	}

	g2 := groceries{
		Item:  "brinjal",
		Price: 20,
	}

	g3 := groceries{
		Item:  "potato",
		Price: 10,
	}

	items = append(items, g1, g2, g3)

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating new file")
	}

	err = t.ExecuteTemplate(newFile, "index.gohtml", items)
	if err != nil {
		log.Fatalln("Error executing template")
	}
}
