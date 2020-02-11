package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

type person struct {
	Name string
	Age  int
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakeArgs(x int) int {
	return x * 2
}

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	p1 := person{
		Name: "Ravikanth",
		Age:  32,
	}

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error in creating the file")
	}

	err = t.ExecuteTemplate(newFile, "index.gohtml", p1)
	if err != nil {
		log.Fatalln("Error in executing the template")
	}

}
