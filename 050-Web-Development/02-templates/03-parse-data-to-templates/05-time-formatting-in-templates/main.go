package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var t *template.Template

var fm = template.FuncMap{
	"format": timeFormatting,
}

func init() {
	t = template.Must(template.New("").Funcs(fm).ParseFiles("foo.gohtml"))
}

func timeFormatting(t time.Time) string {
	return t.Format("2006-01-02")
}

func main() {
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error in creation of the file, failing..")
	}

	err = t.ExecuteTemplate(newFile, "foo.gohtml", time.Now())
	if err != nil {
		log.Fatalln("Error in writing data to the template")
	}
}
