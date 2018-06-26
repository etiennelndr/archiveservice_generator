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

	Services  []Service
	DataTypes []DataType
	Errors    []Error
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
func (s Service) AddOperation(op Operation) {
	s.Operations = append(s.Operations, op)
}

// AddDatatype adds a new data type to the service
func (s Service) AddDatatype(data DataType) {
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

// Message TODO:
type Message struct {
	Name      string
	DataTypes []DataType
}

// DataType TODO:
type DataType struct {
}

// Error TODO:
type Error struct {
}
