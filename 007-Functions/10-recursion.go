/*

Recursion - a function calling itself

*/

package main

import (
	"fmt"
)

func main() {
	number := 5
	n := factorial(number)
	fmt.Printf("The factorial of %d using recursion is %d\n", number, n)

	factorial := 1
	for i := number; i > 0; i-- {
		factorial *= i
	}
	fmt.Printf("The factorial of %d without recursion is %d\n", number, factorial)

}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
