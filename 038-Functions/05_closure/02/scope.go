package main

import (
	"fmt"
)

var x int // package level scope, which also means x = 0

func increment() int {
	x++ // x is accessible here
	return x
}

func main() {
	fmt.Println(increment()) // x = 1
	fmt.Println(increment()) // x = 2
}
