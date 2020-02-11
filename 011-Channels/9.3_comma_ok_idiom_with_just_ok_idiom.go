package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 2)
	go func() {
		channel <- 42
		channel <- 90
		close(channel)
	}()

	v, ok := <-channel
	fmt.Println(v)
	fmt.Println(ok)

	fmt.Println("\n")

	v, ok = <-channel
	fmt.Println(v)
	fmt.Println(ok)

	fmt.Println("\n")

	v, ok = <-channel
	fmt.Println(v)
	fmt.Println(ok)
}
