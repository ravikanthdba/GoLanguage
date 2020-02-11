package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("The number of cores are: ", runtime.NumCPU())
	fmt.Println("The number of Go routines:", runtime.NumGoroutine())

	var counter int = 0;
	var wg sync.WaitGroup;
	const maxCounter = 100;

	wg.Add(maxCounter)

	var mux sync.Mutex;

	for i := 0; i < maxCounter; i++ {
		go func() {
			mux.Lock()
			v := counter
			runtime.Gosched()
			v ++
			counter = v
			mux.Unlock()
			wg.Done()
		}()
		fmt.Println("The number of Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("The number of Go routines:", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}