/*

Functions can be passed in as an argument to another function
Below example provides sum of odd numbers

*/

package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sumNumbers := sum(numbers...)
	fmt.Printf("The overall sum of all the 9 numbers is %d\n", sumNumbers)

	oddNumberSum := oddSum(sum, numbers...)
	fmt.Printf("The sum of all the odd numbers is %d\n", oddNumberSum)

	evenNumberSum := evenSum(sum, numbers...)
	fmt.Printf("The sum of all the even numbers is %d\n", evenNumberSum)

	strings := []string{"Lorem", "Ipsum", "is", "simply", "dummy", "text"}
	vowelNumber := findVowel(sum, strings...)
	fmt.Printf("The number of vowels in the string %s is %d\n", strings, vowelNumber)
}

func sum(numbers ...int) int {
	total := 0
	for number := 0; number < len(numbers); number++ {
		total += numbers[number]
	}
	return total
}

func oddSum(f func(numbers ...int) int, variables ...int) int {
	var yi []int
	for _, variable := range variables {
		if variable%2 != 0 {
			yi = append(yi, variable)
		}
	}

	return f(yi...)
}

func evenSum(f func(numbers ...int) int, variables ...int) int {
	var yi []int
	for _, variable := range variables {
		if variable%2 == 0 {
			yi = append(yi, variable)
		}
	}

	return f(yi...)
}

func findVowel(f func(numbers ...int) int, letter ...string) int {
	var yi []int
	for i := 0; i < len(letter); i++ {
		count := 0
		for j := 0; j < len(letter[i]); j++ {
			if letter[i][j] == 'a' || letter[i][j] == 'e' || letter[i][j] == 'i' || letter[i][j] == 'o' || letter[i][j] == 'u' || letter[i][j] == 'A' || letter[i][j] == 'E' || letter[i][j] == 'I' || letter[i][j] == 'O' || letter[i][j] == 'U' {
				count++
				fmt.Printf("%s %c %d\n", letter[i], letter[i][j], count)
				yi = append(yi, count)
			}
		}
	}
	return f(yi...)
}
