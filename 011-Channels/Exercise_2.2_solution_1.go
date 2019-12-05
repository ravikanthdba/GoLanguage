package main

import (
	"fmt"
)

func main() {
	// c := make(<- chan int); // We are creating a receive only channel and we are sending data to channel in line # 10. Hence this code would not work.
	c := make(chan int) // Making the channel a bi-directional channel will work.

	go func() {
		c <- 42
	}()

	fmt.Println("The value of c is:", <-c)
}
