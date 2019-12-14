package main

import (
	"fmt"
)

func main() {
	var x interface{} = 10
	fmt.Println(x.(int))
}
