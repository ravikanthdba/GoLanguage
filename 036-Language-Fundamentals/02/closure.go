package main

import (
	"fmt"
)

func main() {
	var x int = 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("The value of x is now: ", x)
}
