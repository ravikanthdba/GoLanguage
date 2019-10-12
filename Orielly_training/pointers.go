package main

import (
	"fmt"
)

func double(x int)  {
	fmt.Println(x * x)
}

func doublingNumber(number *int) {
	*number *= 2
	fmt.Println("From the function, the value after doubling is", *number)
}

func main() {
	amount := 6
	double(amount)
	fmt.Println(amount)


	doublingNumber(&amount)
	fmt.Println("From the main function, the value of amount now is ", amount)

	// Pointers
	var x int = 10;
	var y float32 = 20.34;
	var z bool = true;
	fmt.Printf("Address of x: %x\n", &x)
	fmt.Printf("Address of y: %x\n", &y)
	fmt.Printf("Address of z: %x\n", &z)


	// Declare pointer variables
	var myInt int = 67;
	var myIntPointer *int;
	myIntPointer = &myInt
	fmt.Println("Value ", myInt, " is at address: ", myIntPointer)
	fmt.Println("The value at address: ",myIntPointer, " is ", *myIntPointer)

	var myFloat float32 = 10.32;
	var myFloatPointer *float32
	myFloatPointer = &myFloat
	fmt.Println("value ", myFloat, " is at address: ", myFloatPointer)
	fmt.Println("The value at address: ",myFloatPointer, " is ", *myFloatPointer)


	// Update a variable using pointer
	var myvariable int = 10
	fmt.Println("myvariable: ", myvariable)
	var addressmyVariable *int
	addressmyVariable = &myvariable	
	fmt.Println("Address of myvariable", &addressmyVariable)
	fmt.Println("Changing the value of variable now")
	*addressmyVariable = 50
	fmt.Println("value of addressmyVariable", addressmyVariable)
	fmt.Println("value of *addressmyVariable", *addressmyVariable)
	fmt.Println("value of myvariable", myvariable)


	// https://play.golang.org/p/W3yVQrLx7AI - Example from the session
}
