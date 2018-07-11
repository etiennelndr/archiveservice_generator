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

// AbstractTypes adds a new abstract type to this area
func (a Area) AbstractTypes() []Composite {
	var composites []Composite
	for _, c := range a.Composites {
		if c.IsAbstract() {
			composites = append(composites, c)
		}
	}
	return composites
}

// IsAbstractInArea TODO:
func (a Area) IsAbstractInArea(data string) bool {
	for _, c := range a.Composites {
		if c.Name == data {
			return c.IsAbstract()
		}
	}
	return false
}

// Service TODO:
type Service struct {
	Name    string
	Number  string
	Comment string

	Operations   []Operation
	Composites   []Composite
	Enumerations []Enumeration
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

// AbstractTypes adds a new abstract type to this service
func (s Service) AbstractTypes() []Composite {
	var composites []Composite
	for _, c := range s.Composites {
		if c.IsAbstract() {
			composites = append(composites, c)
		}
	}
	return composites
}

// AddOperation adds a new operation to the service
func (s *Service) AddOperation(op Operation) {
	s.Operations = append(s.Operations, op)
}

// AddComposite adds a new composite type to the service
func (s *Service) AddComposite(data Composite) {
	s.Composites = append(s.Composites, data)
}

// AddEnumeration adds a new composite type to the service
func (s *Service) AddEnumeration(data Enumeration) {
	s.Enumerations = append(s.Enumerations, data)
}

// IsAbstractInService TODO:
func (s Service) IsAbstractInService(data string) bool {
	for _, c := range s.Composites {
		if c.Name == data && c.IsAbstract() {
			return true
		}
	}
	return false
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
	Name  string
	Types []Type
}

// AddType TODO:
func (m *Message) AddType(d Type) {
	m.Types = append(m.Types, d)
}

// Type TODO:
type Type struct {
	Name          string
	Comment       string
	ShortFormPart string
	List          string
	Service       string
	Area          string
}

// IsList checks if the type is a list or not
func (t Type) IsList() bool {
	return t.List == "true"
}

// AdaptType is useful to retrieve the real type. Indeed, if
// this type is a list it returns the name of the type + 'List'.
// Otherwise only the name is returned.
func (t Type) AdaptType() string {
	if t.IsList() {
		return t.Name + "List"
	}
	return t.Name
}

// Composite TODO:
type Composite struct {
	Name          string
	ShortFormPart string
	Comment       string
	isAbstract    bool
	// Fields
	Fields []Field
	// Extends
	NameOfTypeToExtend string
	AreaOfTypeToExtend string
}

// NewComposite create a new composite
func NewComposite(name string, comment string, shortFormPart string, nameOfTypeToExtend string, areaOfTypeToExtend string) Composite {
	comp := Composite{
		Name:               name,
		Comment:            comment,
		ShortFormPart:      shortFormPart,
		NameOfTypeToExtend: nameOfTypeToExtend,
		AreaOfTypeToExtend: areaOfTypeToExtend,
	}

	return comp
}

// AddField add a new field to the composite
func (c *Composite) AddField(f Field) {
	c.Fields = append(c.Fields, f)
}

// IsAbstract TODO:
func (c Composite) IsAbstract() bool {
	return c.isAbstract
}

// MakeAbstract TODO:
func (c *Composite) MakeAbstract() {
	c.isAbstract = true
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

// Enumeration TODO:
type Enumeration struct {
	Name          string
	ShortFormPart string
	Comment       string
	Items         []Item
}

// AddItem adds a new item to the enumeration
func (e *Enumeration) AddItem(i Item) {
	e.Items = append(e.Items, i)
}

// Item TODO:
type Item struct {
	Value   string
	NValue  string
	Comment string
}

// Error TODO:
type Error struct {
	Name    string
	Number  string
	Comment string
}
