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

func specs(a interface{}) {
	fmt.Println(a)
	fmt.Printf("The type is %T\n", a)
}

func slicespecs(a ...interface{}) {
	fmt.Println("From the function slicespecs: ", a)
}

func main() {
	c := Cat{}
	d := Dog{}

	sliceAnimals := []interface{}{c, d}
	slicespecs(sliceAnimals)

	specs(c)
	specs(d)
}
