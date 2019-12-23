package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	name := "Ravikanth Garimella"
	body :=
		`<html>
		 <title> Example 2 </title>
         <body>
			<p1> Hello ` + name + ` </p1>
		 </body>
		 </html>`

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("File has not been created, errored out..")
	}
	defer file.Close()

	err = ioutil.WriteFile("index.html", []byte(body), 644)
	if err != nil {
		log.Fatal("ERror writing to the file, failed...")
	}
}
