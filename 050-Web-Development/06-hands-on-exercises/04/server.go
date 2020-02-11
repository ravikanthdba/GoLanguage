package main

import (
	"io"
	"log"
	"net"
)

func main() {
	connection, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatalln("Error in connection: ", err)
	}
	defer connection.Close()


	for {
		conn, err := connection.Accept()
		if err != nil {
			log.Fatalln("Error in accepting connections ", err)
		}

		io.WriteString(conn, "I see you connected\n")
		conn.Close()
	}
}