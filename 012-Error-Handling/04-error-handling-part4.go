package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    f, err := os.Open("names.txt");
    if err != nil {
        fmt.Println(err);
        return;
    }
    defer f.Close();

    bs, err := ioutil.ReadAll(f);
    if err != nil {
        fmt.Println(err);
        return;
    }
    fmt.Println("Actual string is: \t", string(bs));
    fmt.Println("String in bytes is: \t", bs);
}