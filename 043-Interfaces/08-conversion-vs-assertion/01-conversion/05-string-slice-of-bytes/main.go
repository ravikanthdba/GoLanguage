package main

import (
	"fmt"
)

func main() {
	var x string = "HELLO"
	fmt.Println("The string is: ", x)
	fmt.Println("The slice of byte is :", []byte(x))
}
