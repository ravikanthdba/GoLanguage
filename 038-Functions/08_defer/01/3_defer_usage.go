package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src, err := os.Open("source.txt")
	if err != nil {
		fmt.Println("Unable to open the source file, and the error is: ", err)
		return
	}
	defer src.Close()

	fmt.Println(src)

	destination, err := os.Create("target.txt")
	if err != nil {
		fmt.Println("Unable to create the target file, and the error is: ", err)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, src)
	if err != nil {
		fmt.Println("Error copying the file", err)
		return
	}
	fmt.Println("File copy from source to target is done.")
}
