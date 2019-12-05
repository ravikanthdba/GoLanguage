/*

A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument

*/

package main

import (
	"fmt"
)

func main() {
	var num = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := findOddSum(sum, num...)
	fmt.Println(sum)
}

func sum(x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}

func findOddSum(x func(number ...int) int, values ...int) int {
	var numbers []int
	for _, value := range values {
		if value%2 != 0 {
			numbers = append(numbers, value)
		}
	}
	return x(numbers...)
}
