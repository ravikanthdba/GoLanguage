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
		log.Fatalln("Error reading the data from the URL")
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 12)

	for scanner.Scan() {
		n := hashing(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets)
}

func hashing(word string, bucket int) int {
	var sum int
	for _, letter := range word {
		sum += int(letter)
	}
	return sum % bucket
}
