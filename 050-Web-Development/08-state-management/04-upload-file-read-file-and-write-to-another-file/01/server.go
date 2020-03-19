package _1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", fileHandle)
	http.ListenAndServe(":8080", nil)
}

func fileHandle(w http.ResponseWriter, r *http.Request) {
	var s string

	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("fileName")
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}
		fmt.Println("Header: ", header)
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading the content of the file", 500)
			return
		}


		err = ioutil.WriteFile("target/destination.txt", content, 777)
		if err != nil {
			http.Error(w, "Error in writing to new file", 500)
			return
		}
		s = string(content)
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `<form method="post" enctype="multipart/form-data">
	<p> Choose the file to copy </p>
	<input type="file" name="fileName" >
	<input type="submit" >
	</form>
	` + s)
}
