package main

import (
	"fmt"
)

func main() {
	n := Hashbucket("Go", 12)
	fmt.Println(n)
}

// func Hashbucket(word string, bucket int) int {
//     hash := int(word[0])
//     return hash % bucket;
// }

// optionally

func Hashbucket(word string, bucket int32) int32 {
	hash := rune(word[0])
	return hash % bucket
}
