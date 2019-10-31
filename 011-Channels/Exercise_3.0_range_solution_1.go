// Have a code in https://play.golang.org/p/sfyu4Is3FG, we need to use a range over the channel and print the values in the channel.

package main

import (
    "fmt"
)

func main() {
    channel := gen();
    fmt.Printf("The type of c is %T\n", channel);
    fmt.Println("The value of c is: ", channel);
    receive(channel);
    fmt.Println("Exiting the program now.")
}

func receive(c <- chan int) {
    for value := range c {
        fmt.Println("The value is now: ", value);
    }
    fmt.Println("\n")
}

func gen() <- chan int {
    c := make(chan int);

    go func() {
        for i := 0; i < 100; i ++ {
            c <- i;
        }
        close(c);
    }()

    return c;
}