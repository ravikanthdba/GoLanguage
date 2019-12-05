package main

import (
	"fmt"
)

func main() {
	str, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
