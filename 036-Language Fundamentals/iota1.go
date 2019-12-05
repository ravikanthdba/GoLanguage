package main

import (
	"fmt"
)

const (
	a = "a"
	b = iota
	c = iota
)

const (
	d = iota
	e
	f
	g
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
