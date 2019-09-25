package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type Classify struct {
	Works []Work `xml:"works>work"`
} 

type Work struct {
	Title    string `xml:"title,attr"`
	Author   string `xml:"author,attr"`
	Hyr      string `xml:"hyr,attr"`
	Owi      string `xml:"owi,attr"`
}


func main() {
	resp, err := http.Get("http://classify.oclc.org/classify2/Classify?summary=true&title=Huckleberry")		
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var classifiersearchresponse Classify
	err = xml.Unmarshal(body, &classifiersearchresponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(classifiersearchresponse)

}
