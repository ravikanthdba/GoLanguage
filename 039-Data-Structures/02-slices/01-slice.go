package main

import (
	"fmt"
)

func main() {
	var x = []int{2, 3, 4, 5, 6}
	fmt.Println("The value of x is: ", x)
	fmt.Printf("The type of x is %T\n", x)
	fmt.Printf("The address of x is %T\n", &x)

	fmt.Println("Substring: ", x[2:])
	fmt.Println("Pick Element: ", x[2])
	fmt.Println("Check Element availability: ", x[2] == 4)
	fmt.Println("From string without bowling intervals: ", "Ravikanth"[4:7])
	fmt.Println("starting to provided end string: ", "Ravikanth"[:7])
	fmt.Println("Another way of printing eveything: ", "Ravikanth"[:])
}
