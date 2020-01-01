package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error creating a connection: ", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connections: ", err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	data := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			log.Fatalln("No data being sent")
		} else {
			switch {
			case strings.Split(line, " ")[0] == "SET":
				key := strings.Split(line, " ")[1]
				value := strings.Split(line, " ")[2]
				data[key] = value
				fmt.Fprintf(conn, "Added the key: %s and value: %s to the database\n", key, data[key])

			case strings.Split(line, " ")[0] == "GET":
				key := strings.Split(line, " ")[1]
				fmt.Fprintf(conn, "Printing the data for the key: %s is %s\n", key, data[key])

			case strings.Split(line, " ")[0] == "DEL":
				key := strings.Split(line, " ")[1]
				delete(data, key)
				fmt.Fprintf(conn, "Deleted the key %s from the database\n", key)

			case strings.Split(line, " ")[0] == "SHOW":
				fmt.Fprintf(conn, "The data in the database is: \n")
				fmt.Fprintln(conn, data)
			}
		}

	}
	defer conn.Close()
}
