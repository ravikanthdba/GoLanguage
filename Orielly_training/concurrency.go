package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

func responseSize(url string, channel chan int) {
	fmt.Println("Getting ", url)
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	channel <- len(body)
}

func main() {
	start := time.Now()
	channel := make(chan int)
	go responseSize("http://example.com/", channel)
	go responseSize("https://tour.golang.org/", channel)
	go responseSize("https://play.golang.org/", channel)
	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println(time.Since(start))

	//Exercise - https://play.golang.org/p/-42igvcEcEI
}