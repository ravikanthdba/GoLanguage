package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error creating the connection ", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln("Error accepting the connection ", err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	i := 0

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)


		if i == 0 {
			fmt.Println(" method: ", strings.Fields(line)[0])
			fmt.Println(" URI: ", strings.Fields(line)[1])
		}

		if line == "" {
			fmt.Println("Erroring out")
			break
		}

		i ++
	}

	body := `<html> <head></head> <body> HOLY COW THIS IS LOW LEVEL </body> </html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}
