package main

import (
    "fmt"
)

func main() {
    var answer1, answer2, answer3 string;

    fmt.Println("Enter the name:");
    _, err := fmt.Scan(&answer1);
    if err != nil {
        fmt.Println(err);
        return;
    }

    fmt.Println("Enter the favourite game:");
    _, err = fmt.Scan(&answer2);
    if err != nil {
        fmt.Println(err);
        return;
    }


    fmt.Println("Enter the favourite food:");
    _, err = fmt.Scan(&answer3);
    if err != nil {
        fmt.Println(err);
        return;
    }


    
    fmt.Println("\n")
    fmt.Println("The answers you have entered are:\n")

    fmt.Println("Name: \t", answer1);
    fmt.Println("Favourite game: \t", answer2);
    fmt.Println("Favourite food: \t", answer3);
}