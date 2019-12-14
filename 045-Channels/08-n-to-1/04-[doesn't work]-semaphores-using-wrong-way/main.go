package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	fmt.Printf("Channel c is created with address: %d\n", c)
	fmt.Printf("Channel done is created with address: %d\n", done)


	go func() {
		fmt.Println("Launching subprocess1")
		for i := 0; i < 200; i ++ {
			fmt.Printf("Adding value %d to the channel through subprocess1\n", i)
			c <- i
		}
		done <- true
	}()


	go func() {
		fmt.Println("Launching subprocess1")
		for i := 300; i < 500; i ++ {
			fmt.Printf("Adding value %d to the channel through subprocess2\n", i)
			c <- i
		}
		done <- true
	}()


	// We are blocked until done returns true
	fmt.Println("Executing the below lines now:")
	fmt.Printf("Done is now: %v\n", <- done)
	fmt.Printf("Done is now: %v\n", <- done)
	close(c)

	fmt.Println("Executing the range function:")
	for value := range c {
		fmt.Printf("Pulling %d out of the channel c", value)
	}

}
