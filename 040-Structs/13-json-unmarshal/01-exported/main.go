package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	fmt.Println("First name: ", p1.First)
	fmt.Println("Last name: ", p1.Last)
	fmt.Println("Age: ", p1.Age)

	fmt.Println("--------------------------------")

	bs := []byte(`{"First":"Ravikanth","Last": "Garimella", "Age":32}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("First name: ", p1.First)
	fmt.Println("Last name: ", p1.Last)
	fmt.Println("Age: ", p1.Age)
}
