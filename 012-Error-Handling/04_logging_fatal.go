package main

import (
    "log"
    "os"
    "fmt"
    "io/ioutil"
)

func main() {

    defer foo();

    // Creation of log file.

    file, err := os.Create("04_logging_fatal.log");
    if err != nil {
        fmt.Println("Error in creating the log file: ", err);
        return;
    }
    defer file.Close();
    log.SetOutput(file);

    // Reading a file which exists and prints contents.

    file, err = os.Open("names.txt");
    if err != nil {
        log.Fatal("Error in opening the file: ", err);
    }
    defer file.Close();

    bs, err := ioutil.ReadAll(file);
    if err != nil {
        log.Fatal("Error reading the file: ", err);
    }
    fmt.Println("The contents of the file as follows:");
    fmt.Println(string(bs));

    // Reading a file which doesn't exists and prints contents.

    file, err = os.Open("names1.txt");
    if err != nil {
        log.Fatal("Error opening the file: ", err);
    }
    defer file.Close();

    bs, err = ioutil.ReadAll(file);
    fmt.Println("The contents of the file as follows:");
    fmt.Println(string(bs));
}

func foo() {
    fmt.Println("Running foo...");
}