package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var t *template.Template

var functions = template.FuncMap{
	"upperCase":            strings.ToUpper,
	"firstThreeCharacters": firstThreeCharacters,
	"getLDAP":              getLDAP,
}

func firstThreeCharacters(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func getLDAP(s string) string {
	return s[:8]
}

func init() {
	t = template.Must(template.New("").Funcs(functions).ParseFiles("index.gohtml"))
}

type person struct {
	Fname           string
	Lname           string
	FullnameForLDAP string
}

func (p *person) fullname() string {
	return strings.ToLower(string(p.Fname[0]) + p.Lname)
}

func main() {
	people := []person{}

	p1 := person{
		Fname: "Ravikanth",
		Lname: "Garimella",
	}

	p1.FullnameForLDAP = p1.fullname()

	p2 := person{
		Fname: "Nagabhushanam",
		Lname: "Garimella",
	}

	p2.FullnameForLDAP = p2.fullname()

	p3 := person{
		Fname: "Swarna Latha",
		Lname: "Garimella",
	}

	p3.FullnameForLDAP = p3.fullname()

	people = append(people, p1, p2, p3)
	newFile, err := os.Create("input.html")
	if err != nil {
		log.Fatalln("New file hasn't been created..")
	}

	err = t.ExecuteTemplate(newFile, "index.gohtml", people)
}
