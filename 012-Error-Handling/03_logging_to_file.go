package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	logfile, err := os.Create("03_logging_to_file.log")
	if err != nil {
		log.Println("Error creating the log file: ", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	readFile, err := os.Open("name.txt")
	if err != nil {
		log.Println("Error in opening the file: ", err)
		return
	}
	defer readFile.Close()
	bs, err := ioutil.ReadAll(readFile)
	if err != nil {
		log.Println("Error in reading the file: ", err)
		return
	}

	fmt.Println(string(bs))
}
