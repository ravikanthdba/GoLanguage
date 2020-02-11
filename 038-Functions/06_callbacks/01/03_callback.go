package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	greaterThanOne(numbers, func(n int) bool {
		return n > 1
	})
}

func greaterThanOne(numbers []int, callback func(n int) bool) {
	var listNumbers []int

	for _, number := range numbers {
		if callback(number) {
			listNumbers = append(listNumbers, number)
		}
	}

	fmt.Println(listNumbers)
}
