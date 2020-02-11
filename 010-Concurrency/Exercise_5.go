/*
Hands-on exercise #5
Fix the race condition you created in exercise #4 by using package atomic
*/

package main

import (
	"fmt"
	"sync"
	"runtime"
	"sync/atomic"
)

func main() {

	var counter int64;
	var maxCounter int = 1000;

	fmt.Println("The number of CPUs are: \t", runtime.NumCPU());
	fmt.Println("The number of Go routines are: \t", runtime.NumGoroutine())

	var wg sync.WaitGroup;
	wg.Add(maxCounter);

	go func() {
		for i := 0; i < maxCounter; i ++ {
			atomic.AddInt64(&counter, 1);
			runtime.Gosched();
			fmt.Println("Counter: \t", counter);
			fmt.Println("The number of Goroutines are: \t", runtime.NumGoroutine())
			wg.Done()
		}
		fmt.Println("The number of Goroutines are: \t", runtime.NumGoroutine());
	}()

	wg.Wait()
	fmt.Println("The counter value finally is: \t", counter)
}