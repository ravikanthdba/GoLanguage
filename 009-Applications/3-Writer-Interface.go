package main

import (
    "fmt"
    "os"
    "io"
)

func main() {
    fmt.Println("Hello World")
    fmt.Fprintln(os.Stdout, "Welcome to Go Programming Language")
    io.WriteString(os.Stdout, "This is through the io writer\n")
}
