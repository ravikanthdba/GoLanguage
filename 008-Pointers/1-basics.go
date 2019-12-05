/*

Pointers are variables of the address where the variable is stored.

If you want to check at which memory location the variable is stored, we can use "&" operator

The type of address of x in the below example is *int - which is called memory of int
The type of address of y in the below example is *string - which is called memory of string


In the below example, when we assign the address of b with x, the value of b becomes address of x.
If we need to dereference the variable of b, then we give it as *b

In the below example, we have set *b = 99. the value of b is address of x. b = &x - denotes that address of x is assigned to b.
Now - we change the value of *b to 99, which means that the data in the address location *b is modified to 99.

*b has the address location of x, hence the valuee of x is also getting changed.

*/

package main

import (
	"fmt"
)

func main() {

	var x int = 43
	fmt.Println("The value of x is ", x)
	fmt.Println("The address at which x is stored is ", &x)

	fmt.Println("\n\n")

	var y string = "Hello"
	fmt.Println("The value of y is ", y)
	fmt.Println("The address at which y is stored is", &y)

	fmt.Println("\n\n")

	fmt.Printf("The type of variable x is %T\n", x)
	fmt.Printf("The type of variable y is %T\n", y)

	fmt.Println("\n\n")

	fmt.Printf("The type of variable address of x is %T\n", &x)
	fmt.Printf("The type of variable address of y is %T\n", &y)

	fmt.Println("\n\n")

	b := &x
	fmt.Println("The value of b is", b)
	fmt.Println("The value at address of b is ", *b)
	fmt.Printf("The value of type of b is %T\n", b)
	fmt.Printf("The value of type of *b is %T\n", *b)

	fmt.Println("\n\n")

	*b = 99
	fmt.Println("The value of b is", b)
	fmt.Println("The value of *b is ", *b)
	fmt.Println("The value of &b is ", &b)
	fmt.Println("The value of *&b is ", *&b)
	fmt.Println("\n")
	fmt.Println("the value of x is ", x)
	fmt.Println("the value of &x is ", &x)
	// fmt.Println("the value of *x is ", *x);
	fmt.Println("the value of *&x is ", *&x)

}
