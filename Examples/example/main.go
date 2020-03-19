package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	inopsAPI  = "https://inops.corp.linkedin.com/readonly/api/v1/device/"
	alfredAPI = "https://alfred.corp.linkedin.com"
)

type Inops struct {
	ProductionState        string             `json:"production_state"`
	Status                 string             `json:"status"`
	ServiceDetails         *[]service         `json:"service"`
	Hostname               string             `json:"name"`
	OperatingSystemDetails *[]operatingsystem `json:"software"`
}

type service struct {
	Product     string `json:"product"`
	Application string `json:"name"`
}

type operatingsystem struct {
	Kernel      string `json:"kernel"`
	SubCategory string `json:"subcategory"`
	Version     string `json:"version"`
}

func urlRead(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error in Processing the API: ", err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error reading the data: ", err)
	}

	return []byte(data)
}

var inopsResponse []Inops

func getDetailsByPool(poolName string) {
	for {
		log.Println("Running the function")
		inopsAPI := inopsAPI + "?name_range=%" + poolName
		dataResponse := urlRead(inopsAPI)
		err := json.Unmarshal(dataResponse, &inopsResponse)
		if err != nil {
			log.Fatalln("Error in Unmarshaling JSON: ", err)
		}
		fmt.Println(dataResponse)
		time.Sleep(5 * time.Second)
	}
}

func getDetailsByHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hostName := r.URL.Query().Get("hostName")
	fmt.Println(inopsResponse)
	for _, host := range inopsResponse {
		if host.Hostname == hostName {
			fmt.Println(host)
		}
	}

}

// func getDetailsByProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Println(ps)
// 	inopsAPI := inopsAPI + "?name_range=%" + poolName
// }

func main() {
	router := httprouter.New()
	go getDetailsByPool("inops.allocation.pool.lps_search")
	router.GET("/getDetailsByHost", getDetailsByHost)
	// router.GET("/product/:poolName", getDetailsByPool)
	http.ListenAndServe(":8080", router)
}
