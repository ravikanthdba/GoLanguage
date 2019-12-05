package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	fmt.Printf("Channel c has been created with address: %d\n", c)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Launching subroutine1...")
		for i := 0; i < 100; i ++ {
			fmt.Printf("Adding %d to the channel c\n", i)
			c <- i
		}
		wg.Done()
		fmt.Println("Ending subroutine1...")
	}()



	go func() {
		fmt.Println("Launching subroutine2...")
		for i := 300; i < 400; i ++ {
			fmt.Printf("Adding %d to the channel c\n", i)
			c <- i
		}
		wg.Done()
		fmt.Println("Ending subroutine2...")
	}()



	go func() {
		wg.Wait()
		close(c)
	}()


	for value := range c {
		fmt.Printf("Value of c is: %d and the pointer of c is: %d\n", value, c)
	}
}
