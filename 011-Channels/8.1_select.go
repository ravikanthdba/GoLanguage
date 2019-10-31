package main

import (
    "fmt"
)

func main() {
    channel1 := make(chan string);
    channel2 := make(chan string);

    go func() {
        channel1 <- "message";
    }()


    go func() {
        channel2 <- "received";
    }()


    for i:=0; i < 2; i ++ {
        select {
            case v := <- channel1: 
                fmt.Println("Printing the message: ", v);
                
            case v := <- channel2: 
                fmt.Println("Printing the message: ", v);
        }
    }
}