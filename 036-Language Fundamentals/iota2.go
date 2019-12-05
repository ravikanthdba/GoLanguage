package main

import (
	"fmt"
)

const (
	_ = iota
	b = iota
	c = iota * 10
	d = iota * 10
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
