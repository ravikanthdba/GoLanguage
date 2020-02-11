package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/wiki", wiki)
	http.ListenAndServe(":8080", nil)
}

func wiki(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("Wikipedia-logo-en-big.png")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer file.Close()
	fmt.Println(file.Name())
	io.Copy(w, file)
}
