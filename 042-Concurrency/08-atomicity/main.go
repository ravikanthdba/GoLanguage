package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go increment("foo")
	go increment("bar")
	wg.Wait()
	fmt.Println("Final Counter Value: ", counter)
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(10) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter))
	}
	wg.Done()
}
