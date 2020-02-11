/*
   Constant is a variable, that doesn't change
*/

package main

import (
	"fmt"
)

const n = 21

func main() {
	const x = 42
	var a, b int

	fmt.Println("Enter a: ")
	fmt.Scan(&a)

	fmt.Println("Enter b: ")
	fmt.Scan(&b)

	fmt.Println(a+x, b+x)
	fmt.Println(a+n, b+n)
}
