package main

import (
	"fmt"
)

func main() {
	var mapVariable = map[int]string{
		1: "Good Morning",
		2: "Subhodayam",
		3: "Aadab hyderabad",
	}

	fmt.Println(mapVariable)
	delete(mapVariable, 2)

	if val, exists := mapVariable[2]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("Value doesn't exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	fmt.Println(mapVariable)
}
