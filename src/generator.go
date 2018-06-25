/**
 * MIT License
 *
 * Copyright (c) 2018 CNES
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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
		for _, service := range area.Services {
			fmt.Printf("Service name: %v", service.Name)
			fmt.Printf(", Service number: %v", service.Number)
			for _, capabilitySet := range service.Capability {
				capabilitySet.PrintAllOperations()
			}
			fmt.Println("")
		}
	}
}
