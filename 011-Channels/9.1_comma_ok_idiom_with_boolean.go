/*
    with boolean for quit
*/

package main

import (
    "fmt"
)

func main() {
    even_channel := make(chan int);
    odd_channel := make(chan int);
    quit_channel := make(chan bool);

    go send(even_channel, odd_channel, quit_channel);
    receive(even_channel, odd_channel, quit_channel);
    fmt.Println("Exiting")

}

func send(even_channel, odd_channel chan <- int, quit_channel chan <- bool) {
    for i := 0; i < 100; i ++ {
        if i % 2 == 0 {
            even_channel <- i;
        } else {
            odd_channel <- i;
        }
    }
    close(quit_channel);
}

func receive(even_channel, odd_channel <- chan int, quit_channel <- chan bool) {
    for {
        select {
            case v := <- even_channel:
                fmt.Println("From Even channel, the number is: \t", v);
            case v := <- odd_channel:
                fmt.Println("From Odd channel, the number is: \t", v);
            case i, ok := <- quit_channel:
                if !ok {
                    fmt.Println("from comma ok: \t", i, ok);
                    return;
                } else {
                    fmt.Println("from comma ok: \t", i, ok); 
                }
        }
    }
}