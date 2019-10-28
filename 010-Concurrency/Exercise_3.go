/*

Hands-on exercise #3
Using goroutines, create an incrementer program, have a variable to hold the incrementer value
launch a bunch of goroutines
each goroutine should read the incrementer value store it in a new variable yield the processor with runtime.Gosched() increment the new variable
write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition. 
Prove that it is a race condition by using the -race flag

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("The number of CPUs are: \t", runtime.NumCPU());
	fmt.Println("The number of Goroutines are: \t", runtime.NumGoroutine());

	var wg sync.WaitGroup;
	
	var counter int = 0;
	const maxCounter = 1000;

	wg.Add(maxCounter);

	for i := 0; i < maxCounter; i ++ {
		go func() {
			variable := counter;
			runtime.Gosched();
			variable ++;
			counter = variable;
			fmt.Println("The value of Counter is: ", counter)
			wg.Done();
		}()
		fmt.Println("The number of Goroutines: \t", runtime.NumGoroutine());
	}
    wg.Wait();
	fmt.Println("Counter: ", counter);
}