// Implements Synchronization Algorithms

package main

import (
	"fmt"
	"sync"
	"runtime"
	"sync/atomic"
)

func main() {
	fmt.Println("OS Version: \t", runtime.GOOS);
	fmt.Println("Number of CPUs: \t", runtime.NumCPU());
	fmt.Println("Number of Go Routines: \t", runtime.NumGoroutine());

	var counter int64;
	var maxCounter int = 100;

	var wg sync.WaitGroup;
	wg.Add(maxCounter)

	go func() {
		for i := 0; i < maxCounter; i ++ {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter: \t", atomic.LoadInt64(&counter))
			fmt.Println("The number of Go routines are: ", runtime.NumGoroutine())
			wg.Done()
		}
		fmt.Println("The number of Go routines are: ", runtime.NumGoroutine())
	}()
	wg.Wait()
	fmt.Println("Counter: ", counter)
}