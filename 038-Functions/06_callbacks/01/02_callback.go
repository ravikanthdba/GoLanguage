package main

import (
	"fmt"
)

func main() {
	visit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(n int) {
		if n%2 != 0 {
			fmt.Println(n)
		}
	})
}

func visit(n []int, callback func(n int)) {
	for _, number := range n {
		callback(number)
	}
}
