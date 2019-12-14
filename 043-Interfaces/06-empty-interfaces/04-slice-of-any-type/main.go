package main

import (
	"fmt"
)

type Animal struct {
	sound string
}

type Cat struct {
	Animal
	friendly bool
}

type Dog struct {
	Animal
	annoying bool
}

func main() {
	a := Cat{}
	b := Dog{}
	c := Dog{}

	animals := []interface{}{a, b, c}
	fmt.Println(animals)
}
