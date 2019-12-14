package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateSum(a []int) int {
	var sum int

	for _, value := range a {
		sum += value
	}

	return sum
}

func fuelCheck(fuel int) int {
	var listfuel []int

	for fuel >= 0 {
		listfuel = append(listfuel, fuel)
		fuel = (fuel / 3) - 2
	}

	return calculateSum(listfuel[1:])
}

func main() {
	var listOutput []int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	input := bufio.NewScanner(file)
	for input.Scan() {
		value, err := strconv.Atoi(input.Text())
		if err != nil {
			log.Fatal(err)
		}

		listOutput = append(listOutput, fuelCheck(value))
	}

	fmt.Println(calculateSum(listOutput))
}
