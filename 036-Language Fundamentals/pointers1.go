package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println("value of a is: \t", a)
	fmt.Println("address of a is: \t", &a)

	fmt.Println("\n")

	var b *int = &a
	fmt.Println("value of b is: \t", b)
	fmt.Println("address of b is: \t", &b)
	fmt.Println("value of *b is: \t", *b)

}

/*
╰─○ go run pointers1.go
value of a is:   42
address of a is:         0xc0000a8000


value of b is:   0xc0000a8000
address of b is:         0xc00000e018
value of *b is:          42
*/
