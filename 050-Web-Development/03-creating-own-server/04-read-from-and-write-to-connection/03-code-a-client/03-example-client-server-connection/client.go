package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Unable to establish a connection. ", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "I say hi, how are you\n")
}
