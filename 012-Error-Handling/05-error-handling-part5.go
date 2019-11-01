package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    file, err := os.Open("names1.txt");
    if err != nil {
        fmt.Println(err);
        return;
    }

    bs, err := ioutil.ReadAll(file);
    if err != nil {
        fmt.Println(err);
        return;
    }

    fmt.Println(string(bs));
}