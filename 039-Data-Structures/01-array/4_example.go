package main

import (
	"fmt"
)

func main() {
	var x [256]byte
	var y [256]int
	var z [256]string

	for i := 0; i < 256; i++ {
		x[i] = byte(i)
		y[i] = i
		z[i] = string(i)
	}

	fmt.Println(x)
	fmt.Println(y)
	for _, v := range z {
		fmt.Printf("%v - %T - %b\n", v, v, []byte(v))
	}
}
