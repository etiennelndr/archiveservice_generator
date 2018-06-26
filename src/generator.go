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

package src

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/etiennelndr/archiveservice_generator/data"
)

// Generator TODO:
type Generator struct {
	buffer  *bytes.Buffer
	xmlRaw  data.Query
	GenArea Area
}

// OpenAndReadXML TODO:
func (g Generator) OpenAndReadXML(path string) error {
	absPath, _ := filepath.Abs("path")
	xmlFile, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	b, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(b, &g.xmlRaw)
	if err != nil {
		return err
	}

	return nil
}

// RetrieveInformation TODO:
func (g Generator) RetrieveInformation() {
	for _, area := range g.xmlRaw.AreaList {
		for _, service := range area.Services {
			fmt.Printf("Service name: %v", service.Name)
			fmt.Printf(", Service number: %v", service.Number)
			for _, capabilitySet := range service.Capability {
				capabilitySet.PrintAllOperations()
			}
			fmt.Println("")
			fmt.Println("Composite(s):")
			for _, composite := range service.Datas.Composites {
				fmt.Println(composite.Name)
			}
			fmt.Println("\nEnumeration(s):")
			for _, enumeration := range service.Datas.Enumerations {
				fmt.Println(enumeration.Name)
			}
			fmt.Println("")
		}
		fmt.Println("Composite(s):")
		for _, composite := range area.Datas.Composites {
			fmt.Println(composite.Name)
		}
		fmt.Println("\nEnumeration(s):")
		for _, enumeration := range area.Datas.Enumerations {
			fmt.Println(enumeration.Name)
		}
		fmt.Println("\nError(s):")
		for _, err := range area.Errs.Errs {
			fmt.Println(err.Name)
		}
	}
}
