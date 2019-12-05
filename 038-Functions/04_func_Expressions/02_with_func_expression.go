package main

import (
	"fmt"
)

func main() {
	greeting := func() {
		fmt.Println("Hello World")
	}

	greeting()
	fmt.Printf("The type of variable greeting is %T\n", greeting)
}
