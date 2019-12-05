/* Mutex - Mutually Exclusive */

package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	wg.Add(2)
	go increment("foo")
	go increment("bar")
	wg.Wait()
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(100) * time.Millisecond)
		mu.Lock()
		// x := counter - we don't need the commented code now, as at a given time, only one process would access this variable
		// x++
		counter++
		fmt.Printf("%s\n", s)
		fmt.Println("The counter value finally is: ", counter)
		mu.Unlock()
	}
	wg.Done()
}
