package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 33
	}()
	fmt.Println(<-c)
}
