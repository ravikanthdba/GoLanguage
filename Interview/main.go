package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"strings"
	"os"
)


func contains(list []string, text string) bool {
	for values := range list {
		if (text == list[values]) {
			return true
		}
	}
	return false
}

type Hdfssite struct {
	XMLName  xml.Name `xml:"configuration"`
	Text     string   `xml:",chardata"`
	Property struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name"`
		Value string `xml:"value"`
	} `xml:"property"`
} 

func parseXML(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	note := &Hdfssite{}
	_ = xml.Unmarshal([]byte(data), &note)
	return strings.Split(note.Property.Value, ",")
}



func main() {
	mountPointList := parseXML("hdfs-site.xml")

	var mountNumber int
	var disk string
	var mount string
	var dirstructure string

	var mountDirectoryMapping = make([]string)

	fmt.Print("Enter the number of disks you want to mount:")
	fmt.Scanln(&mountNumber)
	
	for i := 0; i < mountNumber; i ++ {
		fmt.Print("Enter the disks you want to mount:")
		fmt.Scanln(&disk)
		fmt.Print("Enter the mount point location:")
		fmt.Scanln(&mount)
		fmt.Print("Enter the directory structure:")
		fmt.Scanln(&dirstructure)

		if contains(mountPointList, mount + dirstructure) == true {
			fmt.Println(mount + dirstructure + "already exists, quitting...")
		} else {
			mountPointList.append(mountPointList, mount + dirstructure)
		}
	}

	fmt.Println(mountPointList)

}
