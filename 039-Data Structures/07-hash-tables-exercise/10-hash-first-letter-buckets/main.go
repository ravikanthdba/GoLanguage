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
		log.Fatalln("Unable to get data from the above link")
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 200)
	for scanner.Scan() {
		n := hashing(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets)
}

func hashing(word string) int32 {
	alphabet := rune(word[0])
	return alphabet
}
