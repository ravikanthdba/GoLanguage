package main

import (
	"fmt"
)

func main() {
	// c := make(chan <- int); // We are creating a send only channel and in line # 13, we are trying to receive data off the channel, hence the code will fail.
	c := make(chan int) // Solution is to change it to bi-directional channel
	go func() {
		c <- 42
	}()

	fmt.Println("The value of c is: ", <-c)
}
