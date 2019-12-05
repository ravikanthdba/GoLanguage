package main

import "fmt"

type foo int

func main() {
	var x foo = 16
	fmt.Println("The value of x is: ", x)
	fmt.Printf("The type of x is %T\n", x)

	var y int = 32
	fmt.Println("The value of y is: ", y)
	fmt.Printf("The type of y is %T\n", y)

	// fmt.Println("The sum of x and y is: ", x+y) This doesn't work as x is of type main.foo and y is of type int
	fmt.Println("The sum of x and y is: ", int(x)+y) // This works as x is a foo, which can be converted to int
	fmt.Println("The sum of x and y is: ", x+foo(y)) // This works as y is an int, which can be converted to foo
}
