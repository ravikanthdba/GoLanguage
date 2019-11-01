package main

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
)

func main() {
    defer foo();

    // file, err := os.Create("05_logging_panic.log");
    // if err != nil {
    //     fmt.Println("Cannot create the log file.");
    //     return;
    // }
    // defer file.Close();
    // log.SetOutput(file);


    file, err := os.Open("names_example.txt");
    if err != nil {
        log.Panicln("Unable to open the file names_example.txt");
    }
    defer file.Close();

    bs, err := ioutil.ReadAll(file);
    if err != nil {
        log.Panicln("Unable to read from the file");
    }
    fmt.Println(string(bs));

}

func foo() {
    fmt.Println("Foo running after Panic has been called.")
}