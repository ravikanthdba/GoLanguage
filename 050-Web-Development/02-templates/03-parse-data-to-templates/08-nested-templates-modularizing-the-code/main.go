package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	Fname string
	Lname string
}

func main() {
	p1 := person{
		Fname: "Ravikanth",
		Lname: "Garimella",
	}

	var err = t.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln("Unable to execute the template")
	}

}
