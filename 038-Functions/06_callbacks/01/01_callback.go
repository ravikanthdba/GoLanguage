/*
   In callback, we pass a function as a parameter to a function
   Example:
       func <function_name> (argument1, variable_name func(int))
*/

package main

import (
	"fmt"
)

/*
   In the below visit function, it passes two arguments:
   1st argument - slice of integers
   2nd argument - "callback" - as the name of the function (it can be anything), and the func(int)
*/

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{10, 2, 30, 4}, func(n int) {
		fmt.Println(n)
	})
}
