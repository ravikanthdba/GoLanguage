package main

import (
	"fmt"
)

func main() {
	x := func(a, b int) int {
		return a + b
	}(2, 3)

	fmt.Println("x = ", x)
}
