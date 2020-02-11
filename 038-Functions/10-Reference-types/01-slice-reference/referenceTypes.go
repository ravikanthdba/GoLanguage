/*
When we pass an array, it passes underlying array address of m to the function.

*/

package main

import (
	"fmt"
)

func main() {
	m := make([]string, 1, 20)
	fmt.Println("The value of m is: ", m)
	fmt.Printf("The type of m is : %T\n", m)
	fmt.Println("The address of m is : ", &m)
	changeMe(m)
	fmt.Println("The value of m is: ", m)
	fmt.Printf("The type of m is : %T\n", m)
	fmt.Println("The address of m is : ", &m)
}

func changeMe(z []string) {
	fmt.Println("The value of z is: ", z)
	fmt.Printf("The type of z is %T\n", z)
	fmt.Println("The address of z is ", &z)
	z[0] = "Todd"
	fmt.Println(z)
	fmt.Printf("The type of z is %T\n", z)
}
