// The code under https://play.golang.org/p/MvL6uamrJP pull the values using select

package main

import (
    "fmt"
)

func main() {
    q := make(chan int);
    c := gen(q);
    receive(c, q);
}

func gen(q chan <- int) <- chan int {
    channel := make(chan int);

    go func() {
        for i := 0; i < 200; i ++ {
            channel <- i;
        }
        q <- 1;
        close(channel);
    }()

    return channel;
}

func receive(c, q <- chan int) {
    for {
            select {
                case v := <- c: 
                    fmt.Println(v);
                case <-q: 
                    return;
                
            }
        }
}