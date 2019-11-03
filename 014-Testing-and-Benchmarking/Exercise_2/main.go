package main

import (
	"fmt"
	"github.com/ravikanthdba/GoLanguage/014-Testing-and-Benchmarking/Exercise_2/quote"
	"github.com/ravikanthdba/GoLanguage/014-Testing-and-Benchmarking/Exercise_2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}