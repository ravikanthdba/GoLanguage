package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	reader := strings.NewReader(`{"First" : "Ravikanth", "Last": "Garimella", "Age": 32}`)
	fmt.Println(reader)
	json.NewDecoder(reader).Decode(&p1)

	fmt.Println("First: ", p1.First)
	fmt.Println("Last: ", p1.Last)
	fmt.Println("Age: ", p1.Age)

}
