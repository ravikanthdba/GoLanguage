/*
   slices and maps are reference types
*/

package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	fmt.Println("The value of m is: ", m)
	fmt.Printf("The type of m is : %T\n", m)
	fmt.Println("The address of m is : ", &m)
	changeMe(m)
	fmt.Println("The value of m is: ", m)
	fmt.Printf("The type of m is : %T\n", m)
	fmt.Println("The address of m is : ", &m)

	fmt.Println("=============================")
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func changeMe(z map[string]int) {
	fmt.Println("The value of z is: ", z)
	fmt.Printf("The type of z is %T\n", z)
	fmt.Println("The address of z is ", &z)
	z["Todd"] = 44
	fmt.Println(z)
	fmt.Printf("The type of z is %T\n", z)
}
