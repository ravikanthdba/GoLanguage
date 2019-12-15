/* This shows different methods of creation of slice
   arrays, slice and maps - are called reference types. They use an underlying data structure.
   Every reference type holds 3 header values - type, length and capacity
*/

package main

import (
	"fmt"
)

func main() {
	var x []string
	fmt.Println(x == nil)
	fmt.Printf("The type of x is %T\n", x)
	fmt.Printf("The address of x is %x\n", &x)
	fmt.Println("The length of x is: %d", len(x))
	fmt.Println("The capacity of x is:", cap(x))

	x = append(x, "Ravikanth")

	fmt.Println(x == nil)
	fmt.Printf("The type of x is %T\n", x)
	fmt.Printf("The address of x is %x\n", &x)
	fmt.Println("The length of x is: %d", len(x))
	fmt.Println("The capacity of x is:", cap(x))

	/* The above example doesn't as yet create a reference type slice. it is just a variable, and assigned to nil */

	y := []string{}
	fmt.Println(y == nil)
	fmt.Printf("The type of y is %T\n", y)
	fmt.Printf("The address of y is %x\n", &y)
	fmt.Println("The length of x is: %d", len(y))
	fmt.Println("The capacity of x is:", cap(y))
	// y[0] = "Ravikanth"
	y = append(y, "Ravikanth")

	fmt.Println(y == nil)
	fmt.Printf("The type of y is %T\n", y)
	fmt.Printf("The address of y is %x\n", &y)
	fmt.Println("The length of x is: %d", len(y))
	fmt.Println("The capacity of x is:", cap(y))
	/* The above is called the shorthand notation, and it creates a slice of the string. But here the length of the slice is not defined. Hence assigning the values, we cannot use slice[index] = variable. We need to use 0. The length of the slice is 0 */

	z := make([]string, 20, 20)
	fmt.Println(y == nil)
	fmt.Printf("The type of y is %T\n", z)
	fmt.Printf("The address of y is %x\n", &z)
	fmt.Println("The length of x is: %d", len(z))
	fmt.Println("The capacity of x is:", cap(y))
	z[0] = "Ravikanth" // Works as we have defined the length and capacity of the slice

	fmt.Println(y == nil)
	fmt.Printf("The type of y is %T\n", z)
	fmt.Printf("The address of y is %x\n", &z)
	fmt.Println("The length of x is: %d", len(z))
	fmt.Println("The capacity of x is:", cap(y))

}
