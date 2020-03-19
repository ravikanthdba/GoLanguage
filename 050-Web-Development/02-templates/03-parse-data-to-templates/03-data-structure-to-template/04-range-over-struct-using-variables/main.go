package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("foo.gohtml"))
}

type person struct {
	Name string
	Age  int
}

func main() {
	var people []person
	p1 := person{
		Name: "Ravikanth",
		Age:  32,
	}

	p2 := person{
		Name: "Nagabhushanam",
		Age:  66,
	}

	p3 := person{
		Name: "Swarna Latha",
		Age:  57,
	}

	p4 := person{
		Name: "Bhargavi",
		Age:  28,
	}

	p5 := person{
		Name: "Viraj",
		Age:  1,
	}

	people = append(people, p1, p2, p3, p4, p5)

	newfile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(newfile, "foo.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}
}
