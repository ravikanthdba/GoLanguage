package main

import (
	"fmt"
)

func main() {
	f := factorial(5)
	fmt.Println(f)
}

func factorial(n int) int {
	var fact int = 1

	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}
