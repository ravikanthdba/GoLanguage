package main

import (
	"fmt"
)

func main() {
	myslice := make([]int, 1)
	fmt.Println("The length of myslice is ", len(myslice))
	fmt.Println("myslice value is: ", myslice)
	myslice[0] = 1
	fmt.Println("The length of myslice is ", len(myslice))
	fmt.Println("myslice value is: ", myslice)
	myslice[0]++
	fmt.Println("The length of myslice is ", len(myslice))
	fmt.Println("myslice value is: ", myslice)
}
