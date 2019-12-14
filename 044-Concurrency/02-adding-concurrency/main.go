package main

import (
	"fmt"
	"time"
)

/*
	3 subprocesses running
	a) main
	b) foo
	c) bar
*/

func main() { // started subprocess1 main in separate thread
	start := time.Now()
	go foo()          // started subprocess2 foo in separate thread
	go bar()          // started subprocess2 bar in separate thread
	end := time.Now() // since no waiting is defined, how much ever possible is executed before the execution of this statement will be printed
	fmt.Println("Elapsed time: ", end.Sub(start))
}

func foo() {
	fmt.Println("Executing foo...")
	for i := 0; i < 45; i++ {
		fmt.Println(i)
	}
	fmt.Println("Completed foo...")
}

func bar() {
	fmt.Println("Executing bar...")
	for i := 0; i < 45; i++ {
		fmt.Println(i)
	}
	fmt.Println("Completed bar...")
}
