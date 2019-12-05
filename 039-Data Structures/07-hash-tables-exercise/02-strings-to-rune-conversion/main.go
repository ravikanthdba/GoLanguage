package main

import (
	"fmt"
)

func main() {
	var x string = "Hello"
	for _, variable := range x {
		fmt.Println(variable)
		fmt.Printf("The type is %T\n", variable)
	}
}
