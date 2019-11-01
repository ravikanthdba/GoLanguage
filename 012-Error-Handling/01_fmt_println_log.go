package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    file, err := os.Open("file1.txt");
    if err != nil {
        fmt.Println("File doesn't exist: ", err);
        return;
    }
    defer file.Close();

    bs, err := ioutil.ReadAll(file);
    if err != nil {
        fmt.Println("Error Reading the file: ", err);
        return;
    }

    fmt.Println(string(bs))
}