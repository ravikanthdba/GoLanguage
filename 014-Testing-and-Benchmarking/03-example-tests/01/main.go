package main

import (
    "fmt"
    "github.com/ravikanthdba/GoLanguage/014-Testing-and-Benchmarking/03-example-tests/01/summation"
)

func main() {
    fmt.Println(summation.Summation(2,3));
    fmt.Println(summation.Summation(2,3, 4, 5, 6, 7, 8));
}