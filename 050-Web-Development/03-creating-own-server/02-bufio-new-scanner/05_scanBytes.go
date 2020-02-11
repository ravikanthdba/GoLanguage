package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	lines := "My name is Ravikanth.\nI work for Linkedin Technologies."
	scanner := bufio.NewScanner(strings.NewReader(lines))
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		lines := scanner.Text()
		fmt.Println(lines)
	}
}
