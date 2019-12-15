package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Printf("The memory address for %d is %v\n", a, &a)
	fmt.Printf("The memory location in integer is %d\n", &a)
}
