package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/etiennelndr/archiveservice_generator/data"
)

func main() {
	fmt.Println("MAL API - Service Generator")

	absPath, _ := filepath.Abs("../archiveservice_generator/XML/ServiceDefCOM.xml")
	xmlFile, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	b, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	var q data.Query
	err = xml.Unmarshal(b, &q)
	if err != nil {
		panic(err)
	}

	for _, area := range q.AreaList {
		for _, service := range area.ServiceList {
			fmt.Printf("Service name: %v", service.Name)
			fmt.Printf(", Service number: %v", service.Number)
			for _, capabilitySet := range service.Capability {
				fmt.Printf(", CapabilitySet: %v", capabilitySet.Number)
				fmt.Printf(", does it have invoke op ? -> %v", capabilitySet.Invoke)
				fmt.Println("")
			}
			fmt.Println("")
			//for
		}
	}
}
