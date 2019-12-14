package main

import (
	"fmt"
)

type Dog struct {
	Age interface{}
}

func main() {
	dog := Dog{}

	dog.Age = 3
	fmt.Printf("The dog is of type %T\n", dog.Age)
	fmt.Println("The value of dog is : ", dog.Age)

	dog.Age = 10.567890
	fmt.Printf("The dog is of type %T\n", dog.Age)
	fmt.Println("The value of dog is : ", dog.Age)

	dog.Age = "not really an age"
	fmt.Printf("The dog is of type %T\n", dog.Age)
	fmt.Println("The value of dog is : ", dog.Age)
}
