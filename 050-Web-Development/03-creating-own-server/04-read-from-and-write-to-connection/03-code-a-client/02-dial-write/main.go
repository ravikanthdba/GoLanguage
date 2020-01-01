package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Unable to connect to localhost 8080: ", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "I say I have dialled you")
}
