package main

import (
    "fmt"
)

func main() {
    c := make(chan int);

    go send(c);
    go receive(c);

    fmt.Println("Program ended..");
}

func send(c chan <- int) {
    c <- 99;
}

func receive(c <- chan int) {
    fmt.Println(<-c);
}