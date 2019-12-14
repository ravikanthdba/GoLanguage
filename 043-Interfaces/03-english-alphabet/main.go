package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/first20hours/google-10000-english/master/google-10000-english-no-swears.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)
	fmt.Println(res.Location())
	fmt.Println(res.Cookies())
}
