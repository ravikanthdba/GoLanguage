package main

import (
	"encoding/json"
	"fmt"
)

/* since struct fields are not exported, after marshalling, the values are not printed */

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{first: "Ravikanth", last: "Garimella", age: 32}
	fmt.Println(p1)
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bs)
}
