package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go increment("foo")
	go increment("bar")
	wg.Wait()
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(10 * time.Millisecond)
		counter = x
		fmt.Println(s, i, "counter: ", counter)
	}
	wg.Done()
}
