package main

import (
	"fmt"
)

func makeGreet() func() {
	return func() {
		fmt.Println("Hello World!!")
	}
}

func main() {
	greet := makeGreet()
	greet()
	fmt.Printf("The type of greet is %T\n", greet)
}
