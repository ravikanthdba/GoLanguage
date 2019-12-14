/*
	For using json.Marshal, it requires the struct name to start with "Capitals", and its attributes also start with "Caps"
*/

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
	p1 := Person{
		First: "Ravikanth",
		Last:  "Garimella",
		Age:   32,
	}

	fmt.Println(p1)

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))

}
