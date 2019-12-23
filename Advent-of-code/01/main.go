package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func output(fuel int) int {
	return (fuel / 3) - 2
}

func totalFuel(fuel []int) int {
	sum := 0
	for _, value := range fuel {
		sum += value
	}
	return sum
}

func main() {
	var array []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		int, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		array = append(array, output(int))
	}

	if scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalFuel(array))
}
