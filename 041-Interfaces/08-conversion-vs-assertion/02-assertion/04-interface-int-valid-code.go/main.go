package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x interface{} = 10
	var falseStatement interface{} = "It is not an integer"

	integer, ok := x.(int)
	switch ok {
	case true:
		fmt.Println(strconv.Itoa(integer) + " is an integer")
	case false:
		str1 := fmt.Sprintf("%v\n", x)
		str2 := fmt.Sprintf("%v\n", falseStatement)
		fmt.Println(str1 + str2)
	default:
		fmt.Println("Something else")
	}
}
