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
		log.Fatalln("Unable to create a connection, failing", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connections, failing", err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if i == 0 {
			mux(conn, line)
		}
		if line == "" {
			break
		}
		i ++
	}
}

func mux(conn net.Conn, line string) {
	method := strings.Split(line, " ")[0]
	fmt.Println("**** METHOD: ", method)
	uri := strings.Split(line, " ")[1]
	fmt.Println("**** URI: ", uri)
	if strings.ToUpper(method) == "GET" && uri == "/" {
		index(conn)
	} else if strings.ToUpper(method) == "GET" && uri == "/home" {
		home(conn)
	} else if strings.ToUpper(method) == "GET" && uri == "/about" {
		about(conn)
	} else if strings.ToUpper(method) == "GET" && uri == "/apply" {
		about(conn)
	} else {
		other(conn)
	}
}

func index(conn net.Conn) {
	body := `<html> <title></title> <body> This is the index page 
	<ul>
	<li> <a href="/about"> About Me </a> </li>
	<li> <a href="/home">  Home Page </a> </li>
	<li> <a href="/other"> Other Page </a> </li>
	<li> <a href="/"> Index Page </a> </li>
	</ul>
	</body> </html>`
	fmt.Fprintf(conn, "HTTP:1/1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", )
	fmt.Fprintf(conn, "Content-Type: %s \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func home(conn net.Conn) {
	body := `<html> <title></title> <body> This is the home page </body> </html>`
	fmt.Fprintf(conn, "HTTP:1/1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", )
	fmt.Fprintf(conn, "Content-Type: %s \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func about(conn net.Conn) {
	body := `<html> <title></title> <body> This is the about page </body> </html>`
	fmt.Fprintf(conn, "HTTP:1/1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", )
	fmt.Fprintf(conn, "Content-Type: %s \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func other(conn net.Conn) {
	body := `<html> <title></title> <body> This is the other page </body> </html>`
	fmt.Fprintf(conn, "HTTP:1/1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", )
	fmt.Fprintf(conn, "Content-Type: %s \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
