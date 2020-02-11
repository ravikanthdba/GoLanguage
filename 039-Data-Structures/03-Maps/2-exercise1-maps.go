package main

import (
	"fmt"
)

func description(z map[int]string) {
	fmt.Println("Is it a nil value? ", (z == nil))
	fmt.Printf("The type is %T\n", z)
	fmt.Println("The value of z is ", z)
}

func main() {
	var x map[int]string // This is never initialized, and hence its nil
	description(x)

	fmt.Println("====================")

	y := map[int]string{}
	description(y)

	fmt.Println("====================")

	z := make(map[int]string)
	description(z)

	fmt.Println("====================")

	w := map[int]string{1: "foo", 2: "bar"}
	description(w)

	_, value := w[2]
	fmt.Println("value is: ", value)

	delete(w, 2)
	fmt.Println(w)

	_, err := w[2]
	fmt.Println(err)
}
