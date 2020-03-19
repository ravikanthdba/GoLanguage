package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

type breakfast struct {
	Name  string
	Style string
	Price float64
}

type lunch struct {
	Name  string
	Style string
	Price float64
}

type dinner struct {
	Name  string
	Style string
	Price float64
}

type menu struct {
	HotelName     string
	BreakfastMenu []breakfast
	LunchMenu     []lunch
	DinnerMenu    []dinner
}

func init() {
	t = template.Must(template.ParseFiles("foo.gohtml"))
}

func main() {

	m1 := menu{
		HotelName: "Saravana Bhavan",
		BreakfastMenu: []breakfast{
			{
				Name:  "Dosa",
				Style: "Indian",
				Price: 20,
			},
			{
				Name:  "Bread",
				Style: "American",
				Price: 25,
			},
			{
				Name:  "Pasta",
				Style: "Italian",
				Price: 50,
			},
		},
		LunchMenu: []lunch{
			{
				Name:  "Thali",
				Style: "Indian",
				Price: 200,
			},
			{
				Name:  "Pizza",
				Style: "Italian",
				Price: 300,
			},
			{
				Name:  "Burger",
				Style: "American",
				Price: 150,
			},
		},
		DinnerMenu: []dinner{
			{
				Name:  "North Indian Thali",
				Style: "Indian",
				Price: 200,
			},
			{
				Name:  "South Indian Thali",
				Style: "Indian",
				Price: 250,
			},
		},
	}

	m2 := menu{
		HotelName: "Udupi Bhavan",
		BreakfastMenu: []breakfast{
			{
				Name:  "Dosa",
				Style: "Indian",
				Price: 15,
			},
			{
				Name:  "Bread",
				Style: "American",
				Price: 50,
			},
			{
				Name:  "Chicken",
				Style: "Thai",
				Price: 100,
			},
		},
		LunchMenu: []lunch{
			{
				Name:  "Thali",
				Style: "Indian",
				Price: 200,
			},
			{
				Name:  "Pasta",
				Style: "Italian",
				Price: 200,
			},
			{
				Name:  "Burger",
				Style: "American",
				Price: 100,
			},
		},
		DinnerMenu: []dinner{
			{
				Name:  "North Indian Thali",
				Style: "Indian",
				Price: 200,
			},
			{
				Name:  "South Indian Thali",
				Style: "Indian",
				Price: 250,
			},
		},
	}

	var menuList []menu
	menuList = append(menuList, m1, m2)

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Unable to create the file")
	}

	err = t.ExecuteTemplate(newFile, "foo.gohtml", menuList)
	if err != nil {
		log.Fatalln("Unable to execute template")
	}

}
