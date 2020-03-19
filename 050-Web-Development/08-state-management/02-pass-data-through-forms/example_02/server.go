package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", statusCode)
	http.HandleFunc("/index", index)

	http.ListenAndServe(":8080", nil)
}

func statusCode(w http.ResponseWriter, h *http.Request) {
	io.WriteString(w, "GOOD")
}

func index(w http.ResponseWriter, h *http.Request) {
	name := h.FormValue("name")
	age := h.FormValue("age")
	sex := h.FormValue("sex")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println("Name: "+ name + "Age: " + age + "Sex: " + sex)
	io.WriteString(w, `<form method="post">
		<input type=text name="name">
		<input type="text" name="age">
		<input type="text" name="sex">
		<input type="submit">
		</form>
	` + "Name: "+ name + "Age: " + age + "Sex: " + sex)
}
