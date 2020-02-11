package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Argument has not been provided, exiting")
	}
	fmt.Println("The first argument is: ", os.Args[0])
	fmt.Println("The second argument is: ", os.Args[1])
	name := os.Args[1]
	fmt.Println("The name entered is: ", name)

	template := `<html>
                 <title></title>
				 <body> 
			     <p1> Hello World, this is: ` + name + ` and how are you doing..
                 </body>
                  </html>`

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error in creating the file...")
	}
	defer file.Close()

	err = ioutil.WriteFile("index.html", []byte(template), 644)
	if err != nil {
		log.Fatal("ERror writing data to the file")
	}
}
