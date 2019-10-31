package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup;

func main() {
    wg.Add(2);
    channel := make(chan int);
    go foo(channel);
    go bar(channel);
    wg.Wait();

}

func foo(c chan <- int) {
    c <- 99;
    wg.Done();
}


func bar(c <- chan int) {
    fmt.Println(<- c);
    wg.Done()
}