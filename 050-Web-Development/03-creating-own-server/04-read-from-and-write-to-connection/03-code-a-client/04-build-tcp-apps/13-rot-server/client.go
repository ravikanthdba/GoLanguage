package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Unable to connect, possibly the server is not running", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "way to go")
}
