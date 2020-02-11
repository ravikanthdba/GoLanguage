/*

Functions - are modular
function - keywords
packages - different files

Syntax:

func (r receiver) <function_name> (parameters) (returns (s)) { ... }

*/

package main

import (
	"fmt"
)

func main() {
	foo()                  // Call functions without arguments
	bar("Ravikanth")       // call function with arguments
	s := woo("MoneyPenny") // call function with arguments
	fmt.Println(s)

	var x int = 10
	var y int = 20
	var sum int
	var difference int
	var multiply int

	sum = addition(x, y)
	difference = subtraction(x, y)
	multiply = product(x, y)

	fmt.Printf("The sum of %d and %d is: %d\n", 10, 20, sum)
	fmt.Printf("The difference of %d and %d is: %d\n", 10, 20, difference)
	fmt.Printf("The product of %d and %d is: %d\n", 10, 20, multiply)

	multiplereturns1, multiplereturns2 := multiplereturns("Ravikanth", "Garimella")

	fmt.Println(multiplereturns1)
	fmt.Println(multiplereturns2)

	comparevariable1, comparevariable2, compareresult := comparision(100, 200)
	fmt.Println(comparevariable1)
	fmt.Println(comparevariable2)
	fmt.Println(compareresult)

	// Summing values using variadic parameters

	variadicParametersSum := variadic(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum of numbers using variadic parameters is ", variadicParametersSum)

	// Unfurling a slice

	Unfurlingslice := []int{2, 3, 4, 5, 6, 7, 8}
	addingNumbers := sumNumbers(Unfurlingslice...)
	fmt.Println(addingNumbers)

	names := []string{"Ravi", "Rahul", "Reena"}
	greeting("Hello", names...)
	greeting("Hello")
}

func foo() {
	fmt.Println("This is a function without parameters")
	fmt.Println("hello world from foo function")
}

/* Everything in Go is pass by value
   What we pass is what we get
*/

func bar(s string) {
	fmt.Println("This is a function with parameters")
	fmt.Println("hello, ", s)
}

func woo(s string) string {
	fmt.Println("This function returns string arguments.")
	return fmt.Sprint("Hello from Woo - ", s)

}

func addition(x int, y int) int {
	fmt.Println("Summing two numbers")
	return x + y
}

func subtraction(x int, y int) int {
	fmt.Println("Subtracting two numbers")
	return x - y
}

func product(x int, y int) int {
	fmt.Println("Multiplying two numbers")
	return x * y
}

// Multiple returns

func multiplereturns(first_name, last_name string) (string, bool) {
	fmt.Println("This is a function with multiple returns")
	first := fmt.Sprint(first_name, " ", last_name)
	boolean_value := false
	return first, boolean_value
}

func comparision(variable1, variable2 int) (int, int, bool) {
	fmt.Println("This is a function with multiple returns")
	first_value := variable1
	second_value := variable2
	compare := (variable1 == variable2)
	return first_value, second_value, compare
}

// Variadic parameters

/*
   If you are unsure, how many elements you are going to pass as input, but if we know the datatype, then we give variadic parameters
   In the below example, we calculate the sum of numbers
*/

func variadic(x ...int) int {
	fmt.Println("The list of x is: ", x)
	fmt.Printf("The datatype of x is %T\n", x)
	sum := 0
	for index, value := range x {
		sum += value
		fmt.Printf("At index %d, we add %d to sum, which is now %d\n", index, value, sum)
	}
	return sum
}

/* Following program sums numbers using unfurling slice */

func sumNumbers(x ...int) int {
	fmt.Println("This function is for sum of variables using unfurling slice...")
	fmt.Println(x)
	sum := 0
	for index, value := range x {
		sum += value
		fmt.Printf("The sum will be %d after adding value at %d position, which is %d\n", sum, index, value)
	}
	return sum
}

/* Variadic parameters has to be at the last. If there is only one parameter, in a function, then it has to be only variadic parameter. */

func greeting(prefix string, names ...string) {
	fmt.Println(prefix, " ", names)
}
