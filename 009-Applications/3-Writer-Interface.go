package main

import (
    "fmt"
    "os"
    "io"
    "time"
)

func main() {
    start_time := time.Now()
    fmt.Println("Hello World")
    fmt.Println("Elapsed: ", time.Since(start_time))
    start_time = time.Now()
    fmt.Fprintln(os.Stdout, "Welcome to Go Programming Language")
    fmt.Println("Elapsed: ", time.Since(start_time))
    start_time = time.Now()
    io.WriteString(os.Stdout, "This is through the io writer\n")
    fmt.Println("Elapsed: ", time.Since(start_time))
}
