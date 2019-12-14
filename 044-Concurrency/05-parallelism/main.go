package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Waitgroup waits for all the goroutines to be completed and collected. Here there are 3 go-routines:
main
foo
bar

Wait groups are counted without the main function, and hence we only consider 2 (foo and bar)
*/
var wg sync.WaitGroup

// sets properties before the start of the program. This system has 8 CPUs, this configuration says the program to leverage all the CPUs.
func init() {
	fmt.Println("The number of CPUs is: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	start := time.Now() // check the time before start of program
	wg.Add(2)           // add the number of go-routines without main. it waits for all the go-routines to complete before proceeding to next step
	go foo()            // Launch subprocess1 to kick off foo
	go bar()            // Launch subprocess2 to kick off bar
	wg.Wait()           // wait for foo anf bar to complete
	end := time.Now()   // end time
	fmt.Println("Elapsed time", end.Sub(start))
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(10 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(10 * time.Millisecond)
	}
	wg.Done()
}
