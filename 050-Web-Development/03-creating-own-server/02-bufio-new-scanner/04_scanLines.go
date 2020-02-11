package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	line := "My name is Ravikanth.\nI work for Linkedin."
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
