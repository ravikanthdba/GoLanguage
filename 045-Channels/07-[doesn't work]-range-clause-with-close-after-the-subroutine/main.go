package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Printf("sending %d on the channel\n", i)
			c <- i // after 0 has been put on the channel
		}
	}()
	fmt.Println("closing the channel")
	close(c) // channel is being closed
	for value := range c {
		fmt.Println("Connecting to the channel")
		fmt.Println(value) // 0 gets printed and post that since the channel is closed it would'nt print
		fmt.Printf("%d is now off the channel\n", value)
	}
}
