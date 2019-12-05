package main

import (
	"fmt"
)

func main() {
	var x = []rune{'H', 'E', 'L', 'L', 'O'}
	fmt.Println("The slice of rune is: ", x)
	fmt.Println("The string of slice of rune is: ", string(x))
}
