package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error creating connection", err)
	}
	defer li.Close()


	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error connecting", err)
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Input passed is: ", line)

			if line == "" {
				fmt.Println("There is no input, breaking the loop")
				break
			}
		}

		fmt.Println("Code entered here")
		io.WriteString(conn, "I see you have connected\n")

		conn.Close()
	}
}
