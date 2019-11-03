package main

import (
    "fmt"
    "github.com/ravikanthdba/GoLanguage/013-Writing_Documentation/Exercise_1/dog"
)

type canine struct {
    name string
    age int
}

func main() {
    dogVariable := canine {
        name: "dog",
        age: dog.Years(10),
    }
    fmt.Println(dogVariable);
}