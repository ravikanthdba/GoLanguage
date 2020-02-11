package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	fmt.Println("The value of c is : ", c)
	fmt.Printf("The format of c is %T\n", c)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
			fmt.Printf("Passing the value %d to the channel\n", c)
			fmt.Println("Code is now stopped")
			fmt.Println("Sleeping")
		}
	}()

	go func() {
		for {
			fmt.Printf("Receiving the value %d to print from the channel\n", c)
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
