package main

import (
	"fmt"
)

var X int = 0 // package level scope

func increment() int {
	X++
	return X
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println("The value of x is : ", X)
}
