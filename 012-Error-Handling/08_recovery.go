package main

import (
    "fmt"
)

func main() {
    f(); // Entering the function f() - Step 1
    fmt.Println("Returning from the fucntion f."); // Step 12 - prints the line.
}

func f() { // Step 9 - receives the panic. The function understands he call is now a panic.
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r); // Step 11 - prints the line.
        }
    }() // This executes at the last line of the function.
    fmt.Println("Calling the function: g"); // Prints the line: Calling the function: g - Step 2
    g(0); // Enters the function g with input 0 - Step 3
    fmt.Println("Returning from the function g"); // This never executes in the program.
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!!"); // - Step 7 - Prints the line "Panicking!!"
        panic(fmt.Sprintf("%v", i)); // - Step 8 - Panics
    } // - Step 4 - checks whether the value of i is greater than 3, and it is not. Hence it has panicked, and it returns to the caller.
    defer fmt.Println("Defering in g", i); // Step 10 - After the panic, the next line runs as it is, and then prints the values, and then goes to the function f, where the panic is waiting.
    fmt.Println("Printing in g: ", i); // - Step 5 - Prints the line.
    g(i + 1); // - Step 6 - Increments by 1, and calls fucntion g(1), once the value of i is 4, and calls the function g, the if condition executes.
}