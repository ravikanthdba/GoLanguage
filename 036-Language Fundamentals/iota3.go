package main

import (
	"fmt"
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // untyped constant
	MB = 1 << (iota * 10) // untyped constant
	GB = 1 << (iota * 10) // untyped constant
	TB = 1 << (iota * 10) // untyped constant
)

const hello string = "World" // typed constant

func main() {
	fmt.Printf("%b\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)

	fmt.Println("\n")
	fmt.Printf("%d\n", MB/1024)
	fmt.Printf("%d\n", GB/1024)

	fmt.Printf("hello is : %s\n", hello)
}
