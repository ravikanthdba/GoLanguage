package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to create a connection, hence failing: ", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln("Unable to accept a connection")
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(r13Server(line))
	}
	defer conn.Close()
	fmt.Println("Connection ended here")
}

func r13Server(word string) string {
	letters := []byte(word)
	var output = make([]byte, len(word))
	for index, letter := range letters {
		if letter <= 109 {
			output[index] = letter + 13
		} else {
			output[index] = letter - 13
		}
	}
	return string(output)
}
