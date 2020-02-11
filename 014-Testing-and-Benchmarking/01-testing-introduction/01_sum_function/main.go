package main

import (
	"fmt"
)

func summingNumbers(variables ...int) int {
	sum := 0

	for _, value := range variables {
		sum += value
	}

	return sum
}

func main() {
	fmt.Printf("%d + %d = %d\n", 5, 6, summingNumbers(5, 6))
	fmt.Printf("%d + %d = %d\n", 10, 12, summingNumbers(10, 12))
	fmt.Printf("%d + %d = %d\n", 100, 200, summingNumbers(100, 200))
}
