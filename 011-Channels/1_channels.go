/* 
    This code doesn't work. The channels block and it needs to be released.
*/

package main

import (
    "fmt"
)

func main() {
        c := make(chan int);
        c <- 33;
        fmt.Println(<-c);
}