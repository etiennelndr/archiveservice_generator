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

package data

import (
	"encoding/xml"
	"fmt"
)

// Query TODO:
type Query struct {
	XMLName  xml.Name `xml:"specification"`
	AreaList []Area   `xml:"area"`
}

// Area TODO:
type Area struct {
	XMLName  xml.Name  `xml:"area"`
	Services []Service `xml:"service"`
	Datas    []Data    `xml:"dataTypes"`
	Errors   []Error   `xml:"errors"`
}

// ------------------- DATA TYPES -------------------

// Data TODO:
type Data struct {
	XMLName      xml.Name      `xml:"dataTypes"`
	Enumerations []Enumeration `xml:"enumeration"`
	Composites   []Composite   `xml:"composite"`
}

// Enumeration TODO:
type Enumeration struct {
	XMLName xml.Name `xml:"enumeration"`
	Items   []Item   `xml:"item"`
}

// Item TODO:
type Item struct {
	XMLName xml.Name `xml:"item"`
	Value   string   `xml:"value,attr"`
	NValue  string   `xml:"nvalue,attr"`
	Comment string   `xml:"comment,attr"`
}

// Composite TODO:
type Composite struct {
	XMLName       xml.Name `xml:"composite"`
	Name          string   `xml:"name,attr"`
	ShortFormPart string   `xml:"shortFormPart,attr"`
	Comment       string   `xml:"comment,attr"`
	Extend        Extends  `xml:"extends"`
}

// Extends TODO:
type Extends struct {
	XMLName      xml.Name `xml:"extends"`
	TypeToExtend Type     `xml:"type"`
}

// Field TODO:
type Field struct {
	XMLName   xml.Name `xml:"field"`
	Name      string   `xml:"name,attr"`
	CanBeNull string   `xml:"canBeNull,attr"`
	Comment   string   `xml:"comment,attr"`
	FieldType Type     `xml:"type"`
}

// Type TODO:
type Type struct {
	XMLName xml.Name `xml:"type"`
	List    string   `xml:"list,attr"`
	Name    string   `xml:"name,attr"`
	Service string   `xml:"service,attr"`
	Area    string   `xml:"area,attr"`
}

// --------------------- ERRORS ---------------------

// Error TODO:
type Error struct {
	XMLName xml.Name `xml:"errors"`
	Name    string   `xml:"name,attr"`
	Number  string   `xml:"number,attr"`
	Comment string   `xml:"comment,attr"`
}

// --------------------- SERVICE --------------------

// Service structure to describe a service
type Service struct {
	XMLName xml.Name `xml:"service"`
	// Name is the name of the service
	Name string `xml:"name,attr"`
	// Number is the number of the service
	Number string `xml:"number,attr"`
	//
	Capability []CapabilitySet `xml:"capabilitySet"`
}

// CapabilitySet TODO:
type CapabilitySet struct {
	XMLName xml.Name `xml:"capabilitySet"`
	Number  string   `xml:"number,attr"`
	// Operations
	// Send
	SendOps []SendIP `xml:"sendIP"`
	// Submit
	SubmitOps []SubmitIP `xml:"submitIP"`
	// Request
	RequestOps []RequestIP `xml:"requestIP"`
	// Invoke
	InvokeOps []InvokeIP `xml:"invokeIP"`
	// Progress
	ProgressOps []ProgressIP `xml:"progressIP"`
	// Publish-Subscribe
	PubSubOps []PubSubIP `xml:"pubsubIP"`
}

// PrintAllOperations TODO:
func (cap CapabilitySet) PrintAllOperations() {
	for _, op := range cap.SendOps {
		op.printOperation()
	}
	for _, op := range cap.SubmitOps {
		op.printOperation()
	}
	for _, op := range cap.RequestOps {
		op.printOperation()
	}
	for _, op := range cap.InvokeOps {
		op.printOperation()
	}
	for _, op := range cap.ProgressOps {
		op.printOperation()
	}
	for _, op := range cap.PubSubOps {
		op.printOperation()
	}
}

// Operation TODO:
type Operation struct {
	Name    string `xml:"name,attr"`
	Number  string `xml:"number,attr"`
	Comment string `xml:"comment,attr"`
}

func (op Operation) printOperation() {
	fmt.Printf("\tOperation -> name = %v, number = %v and comment = %v\n", op.Name, op.Number, op.Comment)
}

// SendIP TODO:
type SendIP struct {
	Operation
	XMLName xml.Name `xml:"sendIP"`
}

// SubmitIP TODO:
type SubmitIP struct {
	Operation
	XMLName xml.Name `xml:"submitIP"`
}

// RequestIP TODO:
type RequestIP struct {
	Operation
	XMLName xml.Name `xml:"requestIP"`
}

// InvokeIP TODO:
type InvokeIP struct {
	Operation
	XMLName xml.Name `xml:"invokeIP"`
}

// ProgressIP TODO:
type ProgressIP struct {
	Operation
	XMLName xml.Name `xml:"progressIP"`
}

// PubSubIP TODO:
type PubSubIP struct {
	Operation
	XMLName xml.Name `xml:"pubsubIP"`
}
