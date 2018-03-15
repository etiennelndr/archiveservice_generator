package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Query struct {
	ServiceList  []Service `xml:"malservice"`
}

type Service struct {
	Name string `xml:"name,attr"`
}

func main() {
	fmt.Println("MAL API - Service Generator")

	absPath, _ := filepath.Abs("../MAL_API_Go_Generator/XML/ServiceDefCOM.xml")
	xmlFile, err := os.Open(absPath)
	if err != nil {
		fmt.Println("Error: can't open file, ", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q Query
	err = xml.Unmarshal(b, &q)
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	for _, name := range q.ServiceList {
		fmt.Printf("Service name: %s\n", name)
	}
}
