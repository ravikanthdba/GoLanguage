package main

import (
	"fmt"
)


func main() {
	c := make(chan int)
	fmt.Printf("Channel created with address: %d and type: %T\n", c, c)
	done := make(chan bool)
	fmt.Printf("Channel created with address: %d and type: %T\n", done, done)


	go func() {
		fmt.Println("Launching subroutine1")
		for i := 0; i <= 100; i ++ {
			fmt.Printf("Adding %d to channel from subroutine1\n", i)
			c <- i
		}
		fmt.Println("Setting done to true in subroutine1")
		done <- true
		fmt.Println("Done is true now in subroutine1")
		fmt.Println("Ending subroutine1")
	}()



	go func() {
		fmt.Println("Launching subroutine2")
		for i := 300; i <= 400; i ++ {
			fmt.Printf("Adding %d to channel from subroutine2\n", i)
			c <- i
		}
		fmt.Println("Setting done to true in subroutine2")
		done <- true
		fmt.Println("Done is true now in subroutine2")
		fmt.Println("Ending subroutine2")
	}()


	go func() {
		fmt.Printf("Done is now: %v\n", <- done)
		fmt.Printf("Done is now: %v\n", <- done)
		fmt.Println("Closing the channel")
		close(c)
	}()


	for value := range c {
		fmt.Println("The value of c is: ", value)
	}
}