package main

import (
	"log"
)

func main() {
	log.Println("This is a regular message.")
	log.Panicln("This is a panic message.")
	log.Fatalln("This is a fatal error.")
	log.Println("This is the end of the function.")
}
