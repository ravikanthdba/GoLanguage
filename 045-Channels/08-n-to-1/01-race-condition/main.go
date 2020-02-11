package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	fmt.Printf("Channel created with address: %d\n", c)
	var wg sync.WaitGroup

	go func() {
		fmt.Println("Launching subroutine1")
		wg.Add(1)
		for i := 0; i < 20; i++ {
			fmt.Printf("Adding %d to the channel c by subprocess1\n", i)
			c <- i
		}
		wg.Done()
		fmt.Println("Ending subroutine1")
	}()

	go func() {
		fmt.Println("Launching subroutine2")
		wg.Add(1)
		for i := 100; i < 200; i++ {
			fmt.Printf("Adding %d to the channel c by subprocess2\n", i)
			c <- i
		}
		wg.Done()
		fmt.Println("Ending subroutine2")
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for value := range c {
		fmt.Println("Value in channel is: ", value)
	}
}
