package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS :", runtime.GOOS);
	fmt.Println("ARCH :", runtime.GOARCH)
	fmt.Println("Number of CPU :", runtime.NumCPU())
	fmt.Println("Number of Goroutines :", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("Number of CPU :", runtime.NumCPU())
	fmt.Println("Number of Goroutines :", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i ++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for j := 0; j < 10; j ++ {
		fmt.Println("bar: ", j)
	}
}