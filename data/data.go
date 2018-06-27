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
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

// Query TODO:
type Query struct {
	XMLName  xml.Name `xml:"specification"`
	AreaList []Area   `xml:"area"`
}

// Area TODO:
type Area struct {
	XMLName xml.Name `xml:"area"`

	Name         string `xml:"name,attr"`
	Number       string `xml:"number,attr"`
	Version      string `xml:"version,attr"`
	Comment      string `xml:"comment,attr"`
	Requirements string `xml:"requirements,attr"`

	Services []Service `xml:"service"`
	Datas    Data      `xml:"dataTypes"`
	Errs     Errors    `xml:"errors"`
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
	XMLName       xml.Name `xml:"enumeration"`
	Name          string   `xml:"name,attr"`
	ShortFormPart string   `xml:"shortFormPart,attr"`
	Comment       string   `xml:"comment,attr"`
	Items         []Item   `xml:"item"`
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
	Fields        []Field  `xml:"field"`
}

// Extends TODO:
type Extends struct {
	XMLName      xml.Name `xml:"extends"`
	TypeToExtend Type     `xml:"type"`
}

// Field TODO:
type Field struct {
	XMLName        xml.Name `xml:"field"`
	Name           string   `xml:"name,attr"`
	FieldCanBeNull string   `xml:"canBeNull,attr"`
	Comment        string   `xml:"comment,attr"`
	FieldType      Type     `xml:"type"`
}

// CanBeNull TODO:
func (f Field) CanBeNull() bool {
	return strings.Contains(f.FieldCanBeNull, "true")
}

// Type TODO:
type Type struct {
	XMLName xml.Name `xml:"type"`
	List    string   `xml:"list,attr"`
	Name    string   `xml:"name,attr"`
	Service string   `xml:"service,attr"`
	Area    string   `xml:"area,attr"`
}

// IsAList TODO:
func (t Type) IsAList() bool {
	return strings.Contains(t.List, "true")
}

// --------------------- ERRORS ---------------------

// Errors TODO:
type Errors struct {
	XMLName xml.Name `xml:"errors"`
	Errs    []Error  `xml:"error"`
}

// Error TODO:
type Error struct {
	XMLName xml.Name `xml:"error"`
	Name    string   `xml:"name,attr"`
	Number  string   `xml:"number,attr"`
	Comment string   `xml:"comment,attr"`
}

// --------------------- SERVICE --------------------

// Service structure to describe a service
type Service struct {
	XMLName xml.Name `xml:"service"`

	Name    string `xml:"name,attr"`
	Number  string `xml:"number,attr"`
	Comment string `xml:"comment,attr"`

	Capability []CapabilitySet `xml:"capabilitySet"`
	Datas      Data            `xml:"dataTypes"`
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

// GenerateOperationHeader TODO;
func (op Operation) GenerateOperationHeader(buf *bytes.Buffer) {
	buf.WriteString("func (p *Pr")
}

// -------------------- Patterns --------------------

// SendIP TODO:
type SendIP struct {
	Operation
	XMLName xml.Name `xml:"sendIP"`
	Message Messages `xml:"messages"`
}

// SubmitIP TODO:
type SubmitIP struct {
	Operation
	XMLName xml.Name `xml:"submitIP"`
	Message Messages `xml:"messages"`
}

// RequestIP TODO:
type RequestIP struct {
	Operation
	XMLName xml.Name `xml:"requestIP"`
	Message Messages `xml:"messages"`
}

// InvokeIP TODO:
type InvokeIP struct {
	Operation
	XMLName xml.Name `xml:"invokeIP"`
	Message Messages `xml:"messages"`
}

// ProgressIP TODO:
type ProgressIP struct {
	Operation
	XMLName xml.Name `xml:"progressIP"`
	Message Messages `xml:"messages"`
}

// PubSubIP TODO:
type PubSubIP struct {
	Operation
	XMLName xml.Name `xml:"pubsubIP"`
	Message Messages `xml:"messages"`
}

// -------------------- Pattern messages --------------------

// Messages TODO:
type Messages struct {
	XMLName       xml.Name             `xml:"messages"`
	Invoke        InvokeMessage        `xml:"invoke"`
	Ack           AckMessage           `xml:"acknowledgement"`
	Response      ResponseMessage      `xml:"response"`
	Progress      ProgressMessage      `xml:"progress"`
	Update        UpdateMessage        `xml:"update"`
	Request       RequestMessage       `xml:"request"`
	Submit        SubmitMessage        `xml:"submit"`
	PublishNotify PublishNotifyMessage `xml:"publishNotify"`
	Send          SendMessage          `xml:"send"`
}

// InvokeMessage TODO:
type InvokeMessage struct {
	XMLName xml.Name `xml:"invoke"`
	Comment string   `xml:"comment,attr"`
	Types   []Type   `xml:"type"`
}

// AckMessage TODO:
type AckMessage struct {
	XMLName xml.Name `xml:"acknowledgement"`
}

// ResponseMessage TODO:
type ResponseMessage struct {
	XMLName xml.Name `xml:"response"`
	Comment string   `xml:"comment,attr"`
	Types   []Type   `xml:"type"`
}

// ProgressMessage TODO:
type ProgressMessage struct {
	XMLName xml.Name `xml:"progress"`
	Comment string   `xml:"comment,attr"`
	Types   []Type   `xml:"type"`
}

// UpdateMessage TODO:
type UpdateMessage struct {
	XMLName xml.Name `xml:"update"`
	Comment string   `xml:"comment,attr"`
	Types   []Type   `xml:"type"`
}

// RequestMessage TODO:
type RequestMessage struct {
	XMLName xml.Name `xml:"request"`
	Comment string   `xml:"comment,attr"`
	Types   []Type   `xml:"type"`
}

// SubmitMessage TODO:
type SubmitMessage struct {
	XMLName xml.Name `xml:"submit"`
	Comment string   `xml:"comment,attr"`
	Types   []Type   `xml:"type"`
}

// PublishNotifyMessage TODO:
type PublishNotifyMessage struct {
	XMLName xml.Name `xml:"publishNotify"`
	Comment string   `xml:"comment,attr"`
	Types   []Type   `xml:"type"`
}

// SendMessage TODO:
type SendMessage struct {
	XMLName xml.Name `xml:"send"`
	Comment string   `xml:"comment,attr"`
	Types   []Type   `xml:"type"`
}
