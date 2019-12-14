package main

import "fmt"

func main() {
	var name interface{} = "Hello"
	str, ok := name.(string)

	if ok {
		fmt.Println(str, " is of type string")
	} else {
		fmt.Println(str, " is not of type string")
	}
}
