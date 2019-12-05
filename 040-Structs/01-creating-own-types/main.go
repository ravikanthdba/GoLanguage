package main

import (
	"fmt"
)

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Println(myAge)
	fmt.Printf("The type is %T\n", myAge)
	// fmt.Printf("The type is %T\n", foo);
}

/* We can create our own types in Go.
struct relies on this concept.
here we create a type foo, which is of type int.
we then create "myAge" of type foo.
since internally foo points to int, it accepts 44 as the number.
checking the type of myAge and value of myAge
*/
