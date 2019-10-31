package main

import (
    "fmt"
)

func main() {
    even_channel := make(chan int);
    odd_channel := make(chan int);
    quit_channel := make(chan int);


    go send(even_channel, odd_channel, quit_channel);
    receive(even_channel, odd_channel, quit_channel);

    fmt.Println("Quitting");
}

func send(even_channel, odd_channel, quit_channel chan <- int) {
    for i := 0; i < 100; i ++ {
        if i % 2 == 0 {
            even_channel <- i;
        } else {
            odd_channel <- i;
        }
    }
    close(quit_channel);
}

func receive(even_channel, odd_channel, quit_channel <- chan int) {
    for {
        select {
            case v := <- even_channel:
                fmt.Println("Even Number: \t", v);
            case v := <- odd_channel:
                fmt.Println("Odd Number: \t", v);
            case i, ok := <- quit_channel:
                if !ok {
                    fmt.Println("From comma ok: \t", i, ok);
                    return;
                } else {
                   fmt.Println("From comma ok: \t", i, ok); 
                }
        }
    }
}