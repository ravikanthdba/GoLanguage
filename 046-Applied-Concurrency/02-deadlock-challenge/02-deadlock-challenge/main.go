package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	for value := range c {
		fmt.Println(value)
	}

}
