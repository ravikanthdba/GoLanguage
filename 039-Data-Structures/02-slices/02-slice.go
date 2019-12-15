package main

import (
	"fmt"
)

func main() {
	var slice = make([]int, 0, 5)
	fmt.Println("value = ", slice)
	fmt.Println("address = ", &slice)
	fmt.Println("length = ", len(slice))
	fmt.Println("capacity = ", cap(slice))

	for i := 0; i < 5; i++ {
		slice = append(slice, i)
	}

	for _, variable := range slice {
		fmt.Printf("value: %d - address: %x - address in int: %d - length: %d - capacity - %d\n", variable, &variable, &variable, len(slice), cap(slice))
	}

	for i := 6; i < 15; i++ {
		slice = append(slice, i)
	}

	for _, variable := range slice {
		fmt.Printf("value: %d - address: %x - address in int: %d - length: %d - capacity - %d\n", variable, &variable, &variable, len(slice), cap(slice))
	}

	var slice1 = []string{"1", "2", "3", "4"}
	var slice2 = []string{"a", "b", "c"}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	// deleting element 'a' from the slice
	slice1 = append(slice1[:4], slice1[5:]...)
	fmt.Println(slice1)

}
