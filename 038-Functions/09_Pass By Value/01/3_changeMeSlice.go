package main

import (
	"fmt"
)

func main() {
	var x = make([]int, 1, 5)
	fmt.Println("Main: Before change x = ", x)
	changeMeSlice(x)
	fmt.Println("Main: After change x = ", x)
}

func changeMeSlice(z []int) {
	fmt.Println("changeMeSlice: Before change z = ", z)
	z[0] = 65536
	fmt.Println("changeMeSlice: After change z = ", z)
}
