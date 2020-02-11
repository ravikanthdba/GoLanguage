package main

import (
	"fmt"
)

func foo(x ...int) {
	if len(x) == 0 {
		fmt.Println("Nothing passed as parameters")
	} else {
		fmt.Println(x)
	}
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	foo(slice...)
	foo()
}
