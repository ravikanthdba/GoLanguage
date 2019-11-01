package main

import (
    "fmt"
)

type custErr struct {
    first string
    last string
}

func (cust custErr) Error() string {
    return fmt.Sprintf("Error for the name: %v and %v", cust.first, cust.last);
}

func main() {
    c1 := custErr {
        first: "Ravikanth", 
        last: "Garimella",
    };

    foo(c1);
}

func foo(e error) {
    fmt.Println("Running the function foo: ", e);
}

