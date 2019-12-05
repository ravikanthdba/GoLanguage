package main

import (
	"fmt"
)

func summation(variables ...int) int {
	var sum int = 0

	for _, variable := range variables {
		sum += variable
	}

	return sum
}

func main() {
	var numberList = []int{2, 3, 4, 5, 6}
	fmt.Println(summation(numberList...))

	numberList = []int{2, 3, 4, 5, 6, 10, 11, 12}
	fmt.Println(summation(numberList...))

}
