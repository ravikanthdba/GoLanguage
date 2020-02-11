package main

import (
	"fmt"
	"runtime"
)

func main() {
	channel := make(chan string)
	fmt.Println("Checking the number of CPUs and Go Routines before kicking off the program:")
	fmt.Println("The number of CPUs are: ", runtime.NumCPU())
	fmt.Println("The number of Go Routines are: ", runtime.NumGoroutine())

	go func() {
		fmt.Println("Checking the number of CPUs and Go Routines in the Goroutine:")
		fmt.Println("The number of CPUs are: ", runtime.NumCPU())
		fmt.Println("The number of Go Routines are: ", runtime.NumGoroutine())
		channel <- "message"
	}()
	fmt.Println("\n")
	fmt.Println("Checking the number of CPUs and Go Routines after the program:")
	fmt.Println("The number of CPUs are: ", runtime.NumCPU())
	fmt.Println("The number of Go Routines are: ", runtime.NumGoroutine())
	fmt.Println("The output is: \t", <-channel)
}
