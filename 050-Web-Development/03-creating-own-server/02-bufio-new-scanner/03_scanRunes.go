package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	line := "My name is Ravikanth\nI work for Linkedin"
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Printf("The format of rune is: %T\n", line)
	}
}
