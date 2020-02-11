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
		log.Fatalln("Error creating the connection", err)
	}
	defer li.Close()


	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error connecting: ", err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			fmt.Println("The lien is blank, Ending request")
			break
		}
	}
}
