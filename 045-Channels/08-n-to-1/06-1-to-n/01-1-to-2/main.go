package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	fmt.Printf("Channel c created with address: %d\n", c)
	fmt.Printf("Channel done created with address: %d\n", done)

	go func() {
		fmt.Println("Subroutine 1 launched to put data onto the channel")
		for i := 0; i < 100; i++ {
			fmt.Printf("Added %d to the channel\n", i)
			c <- i
		}
		fmt.Println("Closing the channel")
		close(c)
	}()

	go func() {
		var subroutine2 []int
		fmt.Println("Subroutine 2 launched to get data from the channel")
		for value := range c {
			subroutine2 = append(subroutine2, value)
			fmt.Printf("Value %d printed from subroutine 2\n", value)
		}
		done <- true
		fmt.Println("Values from subroutine2", subroutine2)
	}()

	go func() {
		var subroutine3 []int
		fmt.Println("Subroutine 3 launched to get data from the channel")
		for value := range c {
			subroutine3 = append(subroutine3, value)
			fmt.Printf("Value %d printed from subroutine 3\n", value)
		}
		done <- true
		fmt.Println("Values from subroutine3", subroutine3)
	}()

	<-done
	<-done
}
