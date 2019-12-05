package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(greet("Ravikanth Garimella"))
	fmt.Println(greets("Ravikanth Garimella"))
}

func greet(name string) ([]string, int) {
	var count int
	var listString []string

	listString = strings.Split(name, " ")
	for i := 0; i < len(listString); i++ {
		count += 1
	}

	return listString, count
}

func greets(name string) (s string) {
	return name[1:]
}
