package main

import (
	"fmt"
)

func wrapper() func() int {
	// x := 0

	var x int // more idiomatic way, which means x = 0

	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper() // func expression
	fmt.Println(increment())
	fmt.Println(increment())

}
