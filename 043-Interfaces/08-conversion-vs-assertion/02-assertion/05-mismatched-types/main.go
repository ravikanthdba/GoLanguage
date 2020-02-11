package main

import "fmt"

func main() {
	var x interface{} = 10
	fmt.Printf("The type of x is %T\n", x) // here it shows that x is of type int, but cannot be used as int. You need to assert it as int before using in any operation
	// fmt.Println(x + 20) // This doesn't work as x is of type interface
	fmt.Println(x.(int) + 20)
}
