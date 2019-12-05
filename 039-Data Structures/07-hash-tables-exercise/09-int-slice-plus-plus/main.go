package main

import (
	"fmt"
)

func main() {
	var x = make([]int, 1)
	fmt.Println(x)
	x[0] = 50
	fmt.Println(x)
	x[0]++
	fmt.Println(x)
}
