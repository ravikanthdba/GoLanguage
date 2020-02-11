package main

import (
	"fmt"
)

func main() {
	var result bool

	result = (true && false) || (false && true) || !(false && false)
	fmt.Println(result)
}
