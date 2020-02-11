package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	response, err := http.NewRequest(http.MethodGet, "https://lva1-etlmon02.corp.linkedin.com:8443/ScannerWeb/lumosFirstSnapshotDetails?cluster=4&srcDatabase=OMS&srcTable=PROMOTIONS2", nil)

	if err != nil {
		fmt.Println("Error calling: ", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("%v", string(data))
	}
}
