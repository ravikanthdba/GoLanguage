package main

import (
    "fmt"
)

func main() {
    even_numbers := make(chan int);
    odd_numbers := make(chan int);
    quit_numbers := make(chan int);

    go send(even_numbers, odd_numbers, quit_numbers);
    receive(even_numbers, odd_numbers, quit_numbers);

    fmt.Println("Finished");
}

func send(even_numbers, odd_numbers, quit_numbers chan <- int) {
    for i := 0; i < 100; i ++ {
        if i % 2 == 0 {
            even_numbers <- i;
        } else {
            odd_numbers <- i;
        }
    }
    close(quit_numbers);
}

func receive(even_numbers, odd_numbers, quit_numbers <- chan int) {
    v, ok := <- even_numbers;
    fmt.Println(v);
    fmt.Println(ok);

    v, ok = <- odd_numbers;
    fmt.Println(v);
    fmt.Println(ok);

}