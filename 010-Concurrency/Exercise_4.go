/*
Hands-on exercise #4
Fix the race condition you created in the previous exercise by using a mutex
it makes sense to remove runtime.Gosched()
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int = 0;
	const maxCounter = 1000;

	 var wg sync.WaitGroup;
	 var mu sync.Mutex;

	fmt.Println("The Number of CPUs: \t", runtime.NumCPU());
	fmt.Println("The Number of Go Routines: \t", runtime.NumGoroutine());

	wg.Add(maxCounter);

	for i := 0; i < maxCounter; i ++ {
		go func() {
			mu.Lock();
			variable := counter;
			variable ++;
			counter = variable;
			wg.Done()
			mu.Unlock()
		}()
		fmt.Println("The Number of Go Routines: \t", runtime.NumGoroutine());
	}
	
	wg.Wait()
	fmt.Println("The counter value now is: ", counter)
}