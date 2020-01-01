package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	line := "My name is Ravikanth.\nI work at Linkedin Technologies"
	scanner := bufio.NewScanner(strings.NewReader(line))
	for scanner.Scan() {
		sentence := scanner.Text()
		fmt.Println(sentence)
	}
}
