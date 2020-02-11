package main

import (
	"fmt"
)

func greet() func() string {
	return func() string {
		return "Hello World!!!"
	}
}

func main() {
	makeGreet := greet()
	fmt.Println(makeGreet())
	fmt.Printf("The type of makeGreet is %T\n", makeGreet)
}
