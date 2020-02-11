package main

import (
	"fmt"
	"runtime"
)


func main() {
	fmt.Println("The Operating System is      :", runtime.GOOS);
	fmt.Println("The Architecture is          :", runtime.GOARCH);
	fmt.Println("The number of goroutines are :", runtime.NumGoroutine());
	fmt.Println("The number of CPU are        :", runtime.NumCPU())
}