package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

var fm = template.FuncMap{
	"even": evenNumber,
}

func evenNumber(n int) bool {
	if n % 2 == 0 {
		return true
	}
	return false
}

func init() {
	t = template.Must(template.New(" ").Funcs(fm).ParseFiles("index.gohtml"))
}

type person struct {
	Name string
	Age int
	IsSecretAgent bool
}

func main() {
	var  people = []person{
			{
				Name:          "Ravikanth",
				Age:           32,
				IsSecretAgent: true,
			},
			{
				Name:          "Bhargavi",
				Age:           28,
				IsSecretAgent: false,
			},
			{
				Name:          "Viraj",
				Age:           1,
				IsSecretAgent: true,
			},
			{
				Name:          "",
				Age:           0,
				IsSecretAgent: false,
			},
		}

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Unable to create the file")
    }

	err = t.ExecuteTemplate(newFile, "index.gohtml", people)
	if err != nil {
		log.Fatalln("Unable to execute the template..")
	}
}
