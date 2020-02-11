package main

import (
	"fmt"
)

type shape struct {
	length  int
	breadth int
}

func switchOnType(x interface{}) {
	switch x.(type) { // x.(type) is called assertion and not conversion. You are checking the type of value passed
	case int:
		fmt.Println("Integer value")

	case string:
		fmt.Println("String Value")

	case shape:
		fmt.Println("Shape Value")

	default:
		fmt.Printf("Any other type and the type is: %T\n", x)
	}
}

func main() {
	var variable1 int = 10
	switchOnType(variable1)

	var hello string = "hello"
	switchOnType(hello)

	x := shape{
		length:  10,
		breadth: 20,
	}
	switchOnType(x)

	y := []int{2, 3, 4, 5}
	switchOnType(y)
}
