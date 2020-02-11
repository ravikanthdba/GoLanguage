package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var t *template.Template

var fm = template.FuncMap{
	"header":    header,
	"bodyTable": bodyTable,
	"trimSpace": trimSpace,
}

func header(s [][]string) []string {
	fmt.Println(s[0])
	return s[0]
}

func bodyTable(s [][]string) [][]string {
	return s[1:]
}

func trimSpace(s string) string {
	return strings.Trim(s, " ")
}

func init() {
	t = template.Must(template.New(" ").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	fileData, err := ioutil.ReadFile("table.csv")
	if err != nil {
		log.Fatalln("Unable to read file..", err)
	}

	r := csv.NewReader(strings.NewReader(string(fileData)))
	record, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Error Reading:", err)
	}
	fmt.Printf("The type of record is %T\n", record)

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Unable to create a file:", err)
	}

	err = t.ExecuteTemplate(newFile, "index.gohtml", record)
	if err != nil {
		log.Fatalln("Unable to execute the template:", err)
	}
}
