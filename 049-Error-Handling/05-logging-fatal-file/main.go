package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("logging.txt")
	if err != nil {
		fmt.Println("Error in creating the logging file")
		return
	}
	fmt.Println(file, ": has been created")
	fmt.Println("Setting output file as : ", file)
	log.SetOutput(file)


	 _, errorLog := os.Open("input.txt")
	 if errorLog != nil {
	 	log.Fatal("Error in opening the file", errorLog)
	 }
}
