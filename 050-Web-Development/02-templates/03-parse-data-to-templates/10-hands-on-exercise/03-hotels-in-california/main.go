package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

type region struct {
	Region string
	Hotels []hotel
}

func init() {
	t = template.Must(template.ParseFiles("foo.gohtml"))
}

func main() {
	var regionList []region

	r1 := region{
		Region: "Northern",
		Hotels: []hotel{
			{
				Name:    "Hotel Taj Krishna",
				Address: "12345 Baker Street",
				City:    "California",
				Zip:     12345,
				Region:  "Northern California",
			},
			{
				Name:    "Hotel Taj Banjara",
				Address: "321 John's Street",
				City:    "California",
				Zip:     12345,
				Region:  "Northern California",
			},
			{
				Name:    "Hotel Taj Falaknuma",
				Address: "112 Indian Street",
				City:    "California",
				Zip:     12345,
				Region:  "Northern California",
			},
			{
				Name:    "Hotel Taj Krishna",
				Address: "11 Pakistan Street",
				City:    "California",
				Zip:     12345,
				Region:  "Northern California",
			},
		},
	}

	r2 := region{
		Region: "Southern",
		Hotels: []hotel{
			{
				Name:    "California Motel",
				Address: "321 Mark's Road",
				City:    "California",
				Zip:     12348,
				Region:  "Southern California",
			},
			{
				Name:    "Hello World",
				Address: "111 Frankenstein's Road",
				City:    "California",
				Zip:     12348,
				Region:  "Southern California",
			},
			{
				Name:    "Come Again!!",
				Address: "1 Worldwide Road",
				City:    "California",
				Zip:     12348,
				Region:  "Southern California",
			},
		},
	}

	r3 := region{
		Region: "Central",
		Hotels: []hotel{
			{
				Name:    "Hotel Linkedin",
				Address: "321 Adam's Street",
				City:    "California",
				Zip:     12345,
				Region:  "Central California",
			},
		},
	}

	regionList = append(regionList, r1, r2, r3)
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Unable to create the file.")
	}

	err = t.ExecuteTemplate(newFile, "foo.gohtml", regionList)
	if err != nil {
		log.Fatalln("Unable to execute template")
	}

}
