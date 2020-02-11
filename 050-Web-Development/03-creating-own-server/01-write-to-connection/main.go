package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error in connection", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "Hello World, this is my first Go Webserver\n")
		fmt.Fprint(conn, "This is through Fprintf\n")
		fmt.Fprint(conn, "This is through second Fprintf")

		conn.Close()
	}
}
