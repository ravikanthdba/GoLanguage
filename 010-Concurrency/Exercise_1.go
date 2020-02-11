/*

Hands-on exercise #1
in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists

*/

package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func function1() {
	for i := 0; i < 20; i ++ {
		fmt.Println(i, "Printing the data from function 1")
	}
	wg.Done()
}

func function2() {
	for i := 0; i < 20; i ++ {
		fmt.Println(i, " Printing the data from function 2")
	}
	wg.Done()
}

func main() {

	fmt.Println("The number of Goroutines are : ", runtime.NumGoroutine())
	fmt.Println("The number of CPUs are : ", runtime.NumCPU())

	wg.Add(2)
	go function1()
	go function2()

	wg.Wait()
	fmt.Println("The number of Goroutines are : ", runtime.NumGoroutine())
	fmt.Println("The number of CPUs are : ", runtime.NumCPU())
}