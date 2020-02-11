package main

import (
	"fmt"
	"strconv"
)

func main() {
	strings := strconv.Itoa(10)
	fmt.Printf("The type of strings is %T\n", strings)
	fmt.Printf("The value of string is %s\n", strings)
}
