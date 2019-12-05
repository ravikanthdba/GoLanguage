/*
	Works for streaming data. Takes a struct and sends out data, which is called encoding.
	like a Writer interface.
	Marshalling - with a string
	encoding - with a streaming.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Ravikanth",
		Last:  "Garimella",
		Age:   32,
	}

	fmt.Println(p1)
	fmt.Println(os.Stdout)
	json.NewEncoder(os.Stdout).Encode(p1)
}
