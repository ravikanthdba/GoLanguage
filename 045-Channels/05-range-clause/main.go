package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i ++ {
			fmt.Printf("value %d has been put on the channel\n", i)
			c <- i
		}
		close(c)
	}()

	for value := range c {
		fmt.Println("WAiting for value on the channel")
		fmt.Printf("Received data on channel\n")
		fmt.Println("The value is: ", value)
	}
}
