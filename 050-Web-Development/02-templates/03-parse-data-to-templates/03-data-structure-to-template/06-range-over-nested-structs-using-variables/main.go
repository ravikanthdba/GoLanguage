package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var t *template.Template

type Customer struct {
	Name string
}

type groceries struct {
	Item     string
	Price    float64
	Quantity float64
}

type ItemCustomer struct {
	Customer
	GroceriesList []groceries
	TotalPrice    float64
}

func (i *ItemCustomer) totalBill() float64 {
	var totalPrice float64

	for _, value := range i.GroceriesList {
		totalPrice += (value.Quantity * value.Price)
	}

	return totalPrice
}

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	CustomerList := []ItemCustomer{}

	i1 := ItemCustomer{
		Customer: Customer{"Ravikanth"},
		GroceriesList: []groceries{
			{
				Item:     "Brinjal",
				Price:    20,
				Quantity: 2,
			},
			{
				Item:     "Potato",
				Price:    10,
				Quantity: 4,
			},
			{
				Item:     "Tomato",
				Price:    30,
				Quantity: 1,
			},
			{
				Item:     "Onion",
				Price:    180,
				Quantity: 1,
			},
			{
				Item:     "Pomegranete",
				Price:    110,
				Quantity: 1,
			},
			{
				Item:     "Banana",
				Price:    20,
				Quantity: 4,
			},
		},
	}

	i1.TotalPrice = i1.totalBill()

	i2 := ItemCustomer{
		Customer: Customer{"Nagabhushanam"},
		GroceriesList: []groceries{
			{
				Item:     "Brinjal",
				Price:    25,
				Quantity: 2,
			},
			{
				Item:     "Potato",
				Price:    10,
				Quantity: 2,
			},
			{
				Item:     "Tomato",
				Price:    34,
				Quantity: 3,
			},
			{
				Item:     "Onion",
				Price:    100,
				Quantity: 1,
			},
			{
				Item:     "Pomegranete",
				Price:    150,
				Quantity: 2,
			},
		},
	}

	i2.TotalPrice = i2.totalBill()

	CustomerList = append(CustomerList, i1, i2)
	fmt.Println(CustomerList)

	newfile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Unable to create the file, please recheck... exiting..")
	}

	err = t.ExecuteTemplate(newfile, "index.gohtml", CustomerList)
	if err != nil {
		log.Fatalln("Unable to write data to template")
	}
}
