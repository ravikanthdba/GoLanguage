package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		fmt.Println("Method is: ", r.Method)
		file, header, err := r.FormFile("fileName")
		fmt.Println("Method is: ", r.Method)
		fmt.Println("file is: ", file)
		fmt.Println("header is: ", header)
		fmt.Println("error is: ", err)
		if err != nil {
			http.Error(w, "Error in capturing the file", 404)
			return
		}
		defer file.Close()
		fmt.Println("The headers are: ", header)

		contents, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Error in fetching the contents of the file", 404)
			return
		}

		s = string(contents)

		newFile, err := os.Create("target/destination.txt")
		defer newFile.Close()
		if err != nil {
			http.Error(w, "Error in creating a target file", 404)
			return
		}
		defer newFile.Close()

		_, err = newFile.Write(contents)
		if err != nil {
			http.Error(w, "Contents not being written to the new file", 404)
		}
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	io.WriteString(w, `<form method="post" enctype="multipart/form-data">
		<p> Choose the file </p>
		<input type="file" name="fileName">
		<input type="submit">
	` + s)
}
