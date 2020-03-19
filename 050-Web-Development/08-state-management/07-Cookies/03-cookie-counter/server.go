package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/setting", cookieSetting)
	http.HandleFunc("/fetchCookie", fetchCookie)
	http.HandleFunc("/abundance", setAbundantCookie)

	http.ListenAndServe(":8080", nil)
}

func cookieSetting(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:       "hello-world",
		Value:      "Cookie Program",
	})
}

func fetchCookie(w http.ResponseWriter, r *http.Request) {
	var cookieCounter int

	cookie, err := r.Cookie("hello-world")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Cookie 1 is: ", cookie)
		cookieCounter += 1
	}

	cookie, err = r.Cookie("code1")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Cookie 2 is: ", cookie)
		cookieCounter += 1
	}

	cookie, err = r.Cookie("code2")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Cookie 3 is: ", cookie)
		cookieCounter += 1
	}

	cookie, err = r.Cookie("code3")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Cookie 4 is: ", cookie)
		cookieCounter += 1
	}


	fmt.Println("The number of cookies are: ", strconv.Itoa(cookieCounter))
}

func setAbundantCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:       "code1",
		Value:      "alpha-beta-gamma",
	})


	http.SetCookie(w, &http.Cookie{
		Name:       "code2",
		Value:      "un-dros-gres",
	})


	http.SetCookie(w, &http.Cookie{
		Name:       "code3",
		Value:      "okati-rendu-moodu",
	})
}



