package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/file", fileOperation)
	http.ListenAndServe(":8080", nil)
}

func fileOperation(w http.ResponseWriter, r *http.Request) {
	var s string

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println("The method is: ", r.Method)
	if r.Method == http.MethodPost {
		fmt.Println(r.FormFile("filename"))
		fileName, fileHeader , err  := r.FormFile("filename")
		if err != nil {
			http.Error(w, "Error From Line23", http.StatusInternalServerError)
			return
		}
		defer fileName.Close()

		fmt.Println("fileName: ", fileName, "fileHeader: ", fileHeader)

		bs, err := ioutil.ReadAll(fileName)
		if err != nil {
			http.Error(w, "Error From Line32", http.StatusInternalServerError)
			return
		}

		s = string(bs)
	}

	io.WriteString(w, `<form method="post" enctype="multipart/form-data">
			<p> Choose the file you want to upload </p>
			<input type="file" name="filename">
			<input type="submit">
	` + s)
}
