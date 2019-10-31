package main

import (
    "fmt"
    "sync"
    "math/rand"
    "time"
)

func main() {
    c1 := make(chan int);
    c2 := make(chan int);

    /* one channel to run the populate function. */
    go populate(c1);

    /* Another channel, hits the fanOutIn function */
    go fanOutIn(c1, c2);

    for v := range c2 {
        fmt.Println(v);
    }

    fmt.Println("About to exit!!")
}

/* This function adds 100 elements to the channel c and then closes the channel */
func populate(c chan int) {
    for i := 0; i < 100; i ++ {
        c <- i;
    }
    close(c);
}


/* This function consists of time taken to consume the variable v1 */
func fanOutIn(c1, c2 chan int) {
    var wg sync.WaitGroup;
    for v := range c1 {
        fmt.Println("The variable is: ", v);
        wg.Add(1);
        go func(v2 int) {
            fmt.Println("Entered another go routine with variable: ", v2);
            c2 <- timeConsumingWork(v2);
            wg.Done();
        }(v)
    }
    wg.Wait();
    close(c2);
}


/* This calculates the time duration for the variable */
func timeConsumingWork(n int) int {
    time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)));
    return n + rand.Intn(1000)
}