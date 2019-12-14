package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom age"`
}

func main() {
	var p1 Person
	fmt.Println("First Name: ", p1.First)
	fmt.Println("Last Name: ", p1.Last)
	fmt.Println("Age: ", p1.Age)
	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println()
	bs := []byte(`{"First": "Ravikanth", "Last": "Garimella", "wisdom age" : 32}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("First Name: ", p1.First)
	fmt.Println("Last Name: ", p1.Last)
	fmt.Println("Age: ", p1.Age)
}
