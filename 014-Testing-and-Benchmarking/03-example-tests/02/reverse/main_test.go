package reverse

import (
    "fmt"
)

func ExampleReverseString() {
    fmt.Println(ReverseString("Hello"));
    // Output:
    // olleH
}

func ExampleReverseInteger() {
    fmt.Println(ReverseInteger(98761));
    // Output:
    // 16789
}