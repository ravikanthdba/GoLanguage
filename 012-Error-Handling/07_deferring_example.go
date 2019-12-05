package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This program copies the file from source to target. This shows 'defer' as an example.")
	source_file, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error in opening the file")
		return
	}
	defer source_file.Close() // Executes this command at the end of the program.

	source_data, err := ioutil.ReadAll(source_file)
	if err != nil {
		fmt.Println("Failed to read data from the file.")
		return
	}

	target_file, err := os.Create("07_target.txt")
	if err != nil {
		fmt.Println("Error in creating the targaet file")
		return
	}
	defer target_file.Close() // This executes at the second last before closing the program

	n2, err := target_file.Write(source_data)
	if err != nil {
		fmt.Println("Writing to target has failed!!")
		return
	}

	fmt.Printf("Written %d bytes\n", n2)
}
