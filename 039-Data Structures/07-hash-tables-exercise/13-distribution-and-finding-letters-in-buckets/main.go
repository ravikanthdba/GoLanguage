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
		log.Fatalln("Error pulling data from the link")
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12)
	for scanner.Scan() {
		word := scanner.Text()
		n := hashing(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	for id, bucket := range buckets {
		fmt.Println("Bucket ID:", id, " --- > ", len(bucket))
	}

	fmt.Println(buckets[0])
}

func hashing(word string, bucket int) int {
	var sum int
	for _, letter := range word {
		sum += int(letter)
	}
	return sum % bucket
}
