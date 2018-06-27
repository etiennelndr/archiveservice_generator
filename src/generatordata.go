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

// Area TODO:
type Area struct {
	Name         string
	Number       string
	Version      string
	Comment      string
	Requirements string

	Services   []Service
	Composites []Composite
	Errors     []Error
}

// CreateArea creates a new area and returns it
func CreateArea(name string, number string, version string, comment string, requirements string) Area {
	area := Area{
		Name:         name,
		Number:       number,
		Version:      version,
		Comment:      comment,
		Requirements: requirements,
	}
	return area
}

// AddComposite TODO:
func (a *Area) AddComposite(c Composite) {
	a.Composites = append(a.Composites, c)
}

// AddError TODO:
func (a *Area) AddError(e Error) {
	a.Errors = append(a.Errors, e)
}

// AddService TODO:
func (a *Area) AddService(s Service) {
	a.Services = append(a.Services, s)
}

// Service TODO:
type Service struct {
	Name    string
	Number  string
	Comment string

	Operations []Operation
	DataTypes  []DataType
}

// CreateService creates a new service and returns it
func CreateService(name string, number string, comment string) Service {
	service := Service{
		Name:    name,
		Number:  number,
		Comment: comment,
	}
	return service
}

// AddOperation adds a new operation to the service
func (s *Service) AddOperation(op Operation) {
	s.Operations = append(s.Operations, op)
}

// AddDatatype adds a new data type to the service
func (s *Service) AddDatatype(data DataType) {
	s.DataTypes = append(s.DataTypes, data)
}

// Operation TODO:
type Operation struct {
	Name    string
	Number  string
	Comment string

	Pattern PatternInteraction
}

// PatternInteraction TODO:
type PatternInteraction struct {
	Name     string
	Messages []Message
}

// AddMessage TODO:
func (p *PatternInteraction) AddMessage(m Message) {
	p.Messages = append(p.Messages, m)
}

// Message TODO:
type Message struct {
	Name      string
	DataTypes []DataType
}

// AddDataType TODO:
func (m *Message) AddDataType(d DataType) {
	m.DataTypes = append(m.DataTypes, d)
}

// DataType TODO:
type DataType struct {
	DataName      string
	DataComment   string
	ShortFormPart string
	List          string
	Service       string
	Area          string
}

// Composite TODO:
type Composite struct {
	DataType
	// Fields
	Fields []Field
	// Extends
	NameOfTypeToExtend string
	AreaOfTypeToExtend string
}

// NewComposite create a new composite
func NewComposite(dataName string, dataComment string, shortFormPart string, nameOfTypeToExtend string, areaOfTypeToExtend string) Composite {
	comp := Composite{
		DataType: DataType{
			DataName:      dataName,
			DataComment:   dataComment,
			ShortFormPart: shortFormPart,
		},
		NameOfTypeToExtend: nameOfTypeToExtend,
		AreaOfTypeToExtend: areaOfTypeToExtend,
	}

	return comp
}

// AddField add a new field to the composite
func (c *Composite) AddField(f Field) {
	c.Fields = append(c.Fields, f)
}

// Field TODO:
type Field struct {
	Name      string
	CanBeNull string
	Comment   string
	// Type
	TypeName string
	TypeArea string
}

// Error TODO:
type Error struct {
	Name    string
	Number  string
	Comment string
}
