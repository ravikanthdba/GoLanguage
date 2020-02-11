package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var t *template.Template

var functions = template.FuncMap{
	"double": doubling,
	"square": squaring,
	"sqroot": sqroot,
}

func doubling(n float64) float64 {
	return n * 2
}

func squaring(n float64) float64 {
	return math.Pow(n, 2)
}

func sqroot(n float64) float64 {
	return math.Sqrt(n)
}

func init() {
	t = template.Must(template.New("").Funcs(functions).ParseFiles("index.gohtml"))
}

func main() {
	var number float64 = 10

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating the file..")
	}

	err = t.ExecuteTemplate(newFile, "index.gohtml", number)
	if err != nil {
		log.Fatalln("Unable to write data to the template")
	}
}
