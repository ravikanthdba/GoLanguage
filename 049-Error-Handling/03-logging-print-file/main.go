/* This program creates a file and logs errors in a file */

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("logfile.txt")
	if err != nil {
		fmt.Println("error in creating the logfile")
		return
	}

	fmt.Println("file: ", file, " has been created")
	log.SetOutput(file)

	_, errlog := os.Open("text.txt")
	if errlog != nil {
		log.Println("error in opening the file: text.txt")
	}

}
