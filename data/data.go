package data

import "encoding/xml"

// Query TODO:
type Query struct {
	XMLName  xml.Name `xml:"specification"`
	AreaList []Area   `xml:"area"`
}

// Area TODO:
type Area struct {
	XMLName     xml.Name  `xml:"area"`
	ServiceList []Service `xml:"service"`
}

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
	Invoke  InvokeIP `xml:"invokeIP"`
}

// Operation TODO:
type Operation struct {
	Name    string `xml:"name,attr"`
	Number  string `xml:"number,attr"`
	Comment string `xml:"comment,attr"`
}

// InvokeIP TODO:
type InvokeIP struct {
	Operation
	XMLName xml.Name `xml:"invokeIP"`
}

// RequestIP TODO:
type RequestIP struct {
	Operation
	XMLName xml.Name `xml:"requestIP"`
}
