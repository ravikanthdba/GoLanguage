package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println("WAiting to receive baton")
	fmt.Println("received baton")
	fmt.Println(<-c)
}
