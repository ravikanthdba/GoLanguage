package main

import (
	"fmt"
)

func foo(z *int) {
	fmt.Println("The value of z before change is: ", z)
	fmt.Println("The address of z before change is: ", &z)
	*z = 100
	fmt.Println("The value of z after change is: ", z)
	fmt.Println("The address of z after change is: ", &z)
}

func main() {
	x := 66
	fmt.Println("From main: The value of x is: ", x)
	fmt.Println("From main: The address of x is: ", &x)
	fmt.Println("Calling foo...")
	foo(&x)
	fmt.Println("After calling foo, from main: The value of x is: ", x)
	fmt.Println("After calling foo, from main: The address of x is: ", &x)
}
