/*
   Everything in Go is "pass by value".
   We do not pass anything by reference into the function.
   Even if we pass a pointer variable, it is a value and not reference
*/

package main

import (
	"fmt"
)

func main() {
	var x int = 42
	fmt.Println("Before changing, the value is: ", x)
	changeMe(&x)
	fmt.Println("After changing, the value is: ", x)
}

func changeMe(z *int) {
	fmt.Println("From the function changeMe, The value of z is: ", z)
	fmt.Println("From the function changeMe, The value of z is: ", *z)
	*z = 12
	fmt.Println("From the function changeMe, The value of z is: ", z)
	fmt.Println("From the function changeMe, The value of z is: ", *z)
}
