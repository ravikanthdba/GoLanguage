package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("files.txt")
	if err != nil {
		log.Println("Error opening the file: ", err)
		return
	}
	defer file.Close()

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading from the file: ", err)
		return
	}

	fmt.Println(string(bs))
}
