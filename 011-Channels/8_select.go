/*
    select is used for pulling off of multiple channels.
*/

package main

import (
    "fmt"
)

func main() {
    odd_channel := make(chan int);
    even_channel := make(chan int);
    quit_channel := make(chan int);
    go send(odd_channel, even_channel, quit_channel);
    receive(odd_channel, even_channel, quit_channel);

    fmt.Println("Exiting")
}

func send(odd_channel, even_channel, quit_channel chan <- int) {
    for i := 0; i < 100; i++ {
        if (i % 2 == 0) {
            even_channel <- i;
        } else {
            odd_channel <- i;
        }
    }
    quit_channel <- 0;
    close(even_channel);
    close(odd_channel);
    close(quit_channel);
}


func receive(odd_channel, even_channel, quit_channel <- chan int) {
    for {
        select {
            case v := <- even_channel:
                fmt.Println("From the Even channel: \t", v);
            case v := <- odd_channel:
                fmt.Println("From the Odd channel : \t",  v);
            case v := <- quit_channel:
                fmt.Println("From the Quit channel: \t",  v);
                return;
        }
    }
}


