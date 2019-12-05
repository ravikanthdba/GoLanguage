package main

import (
	"fmt"
)

func main() {
	var x interface{} = 10
	string, ok := x.(string)

	if ok {
		fmt.Println(string, " is a type string")
	} else {
		fmt.Println("ok value is: ", ok)
	}
}
