package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln("Error reading from the link")
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanWords)
	bucket := make([]int, 12)
	for scanner.Scan() {
		n := hashing(scanner.Text(), 12)
		bucket[n]++
	}

	fmt.Println(bucket)
}

func hashing(word string, bucket int32) int32 {
	return rune(word[0]) % 12
}
