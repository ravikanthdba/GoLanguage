package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	var wg sync.WaitGroup;
	const gs = 100;
	wg.Add(gs)

	for i := 0; i < gs; i ++ {
		go func() {

			v := counter
			runtime.Gosched()
			v ++;
			counter = v

			wg.Done()
			fmt.Println("GoRoutines: ", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}