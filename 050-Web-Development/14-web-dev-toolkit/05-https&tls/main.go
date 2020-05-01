package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServeTLS(":10430", "cert.pem", "key.pem", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is a TLS request")
}
