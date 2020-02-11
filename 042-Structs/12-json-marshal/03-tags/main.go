package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom source"`
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
	fmt.Println(string(bs))
}
