package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10) // buffered channel, it can store 10 items in the channel at one shot.

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Printf("Sending %d to the channel\n", i)
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Printf("Printing %d from the channel\n", c)
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)

}
