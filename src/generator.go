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
func (g *Generator) OpenAndReadXML(path string) error {
	absPath, _ := filepath.Abs(path)
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
func (g *Generator) RetrieveInformation() {
	// Retrieve the area and its datas
	for _, area := range g.xmlRaw.AreaList {
		// Firstly, retrieve the Name, Number, Version, Comment and Requirements of the area
		g.GenArea.Name = area.Name
		g.GenArea.Number = area.Number
		g.GenArea.Version = area.Version
		g.GenArea.Comment = area.Comment
		g.GenArea.Requirements = area.Requirements

		// Create the composites of this area
		for _, composite := range area.Datas.Composites {
			comp := NewComposite(composite.Name,
				composite.Comment,
				composite.ShortFormPart,
				composite.Extend.TypeToExtend.Name,
				composite.Extend.TypeToExtend.Area)
			for _, field := range composite.Fields {
				f := Field{
					CanBeNull: field.FieldCanBeNull,
					Comment:   field.Comment,
					Name:      field.Name,
					TypeArea:  field.FieldType.Area,
					TypeName:  field.FieldType.Name,
				}
				// Now add this new field to the composite
				comp.AddField(f)
			}
			// Then add it to the area
			g.GenArea.AddComposite(comp)
		}

		// Create the errors of this area
		for _, err := range area.Errs.Errs {
			e := Error{
				Comment: err.Comment,
				Name:    err.Name,
				Number:  err.Number,
			}
			// Then add it to the area
			g.GenArea.AddError(e)
		}
	}

	// Retrieve the service and its operations
	for _, area := range g.xmlRaw.AreaList {
		for _, service := range area.Services {
			s := Service{
				Comment: service.Comment,
				Name:    service.Name,
				Number:  service.Number,
			}

			for _, capabilitySet := range service.Capability {
				for _, op := range capabilitySet.SendOps {
					AddSendOperation(&s, op)
				}
				for _, op := range capabilitySet.SubmitOps {
					AddSubmitOperation(&s, op)
				}
				for _, op := range capabilitySet.RequestOps {
					AddRequestOperation(&s, op)
				}
				for _, op := range capabilitySet.InvokeOps {
					AddInvokeOPeration(&s, op)
				}
				for _, op := range capabilitySet.ProgressOps {
					AddProgressOperation(&s, op)
				}
				for _, op := range capabilitySet.PubSubOps {
					AddPubSubOperation(&s, op)
				}
			}
			g.GenArea.AddService(s)
		}
	}
}

// AddSendOperation TODO:
func AddSendOperation(s *Service, operation data.SendIP) {
	op := Operation{
		Comment: operation.Comment,
		Name:    operation.Name,
		Number:  operation.Number,
		Pattern: PatternInteraction{
			Name: "send",
		},
	}

	// Send Message
	send := Message{
		Name: "send",
	}
	for _, t := range operation.Message.Send.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		send.AddDataType(data)
	}
	op.Pattern.AddMessage(send)

	// Add this new operation to the service
	s.AddOperation(op)
}

// AddSubmitOperation TODO:
func AddSubmitOperation(s *Service, operation data.SubmitIP) {
	op := Operation{
		Comment: operation.Comment,
		Name:    operation.Name,
		Number:  operation.Number,
		Pattern: PatternInteraction{
			Name: "submit",
		},
	}

	// Submit Message
	submit := Message{
		Name: "submit",
	}
	for _, t := range operation.Message.Submit.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		submit.AddDataType(data)
	}
	op.Pattern.AddMessage(submit)

	// Ack Message
	ack := Message{
		Name: "ack",
	}
	op.Pattern.AddMessage(ack)

	// Add this new operation to the service
	s.AddOperation(op)
}

// AddRequestOperation TODO:
func AddRequestOperation(s *Service, operation data.RequestIP) {
	op := Operation{
		Comment: operation.Comment,
		Name:    operation.Name,
		Number:  operation.Number,
		Pattern: PatternInteraction{
			Name: "request",
		},
	}

	// Request Message
	request := Message{
		Name: "request",
	}
	for _, t := range operation.Message.Request.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		request.AddDataType(data)
	}
	op.Pattern.AddMessage(request)

	// Response Message
	response := Message{
		Name: "response",
	}
	for _, t := range operation.Message.Response.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		response.AddDataType(data)
	}
	op.Pattern.AddMessage(response)

	// Add this new operation to the service
	s.AddOperation(op)
}

// AddInvokeOPeration TODO:
func AddInvokeOPeration(s *Service, operation data.InvokeIP) {
	op := Operation{
		Comment: operation.Comment,
		Name:    operation.Name,
		Number:  operation.Number,
		Pattern: PatternInteraction{
			Name: "invoke",
		},
	}

	// Invoke Message
	invoke := Message{
		Name: "invoke",
	}
	for _, t := range operation.Message.Invoke.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		invoke.AddDataType(data)
	}
	op.Pattern.AddMessage(invoke)

	// Ack Message
	ack := Message{
		Name: "ack",
	}
	op.Pattern.AddMessage(ack)

	// Response Message
	response := Message{
		Name: "response",
	}
	for _, t := range operation.Message.Response.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		response.AddDataType(data)
	}
	op.Pattern.AddMessage(response)

	// Add this new operation to the service
	s.AddOperation(op)
}

// AddProgressOperation TODO:
func AddProgressOperation(s *Service, operation data.ProgressIP) {
	op := Operation{
		Comment: operation.Comment,
		Name:    operation.Name,
		Number:  operation.Number,
		Pattern: PatternInteraction{
			Name: "progress",
		},
	}

	// Invoke Message
	progress := Message{
		Name: "progress",
	}
	for _, t := range operation.Message.Progress.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		progress.AddDataType(data)
	}
	op.Pattern.AddMessage(progress)

	// Invoke Message
	update := Message{
		Name: "update",
	}
	for _, t := range operation.Message.Update.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		update.AddDataType(data)
	}
	op.Pattern.AddMessage(update)

	// Response Message
	response := Message{
		Name: "response",
	}
	for _, t := range operation.Message.Response.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		response.AddDataType(data)
	}
	op.Pattern.AddMessage(response)

	// Add this new operation to the service
	s.AddOperation(op)
}

// AddPubSubOperation TODO:
func AddPubSubOperation(s *Service, operation data.PubSubIP) {
	op := Operation{
		Comment: operation.Comment,
		Name:    operation.Name,
		Number:  operation.Number,
		Pattern: PatternInteraction{
			Name: "pubsub",
		},
	}

	// PublishNotify Message
	publishNotify := Message{
		Name: "publishNotify",
	}
	for _, t := range operation.Message.PublishNotify.Types {
		data := DataType{
			Area:     t.Area,
			DataName: t.Name,
			List:     t.List,
			Service:  t.Service,
		}
		publishNotify.AddDataType(data)
	}
	op.Pattern.AddMessage(publishNotify)

	// Add this new operation to the service
	s.AddOperation(op)
}
