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
	"strings"

	"github.com/etiennelndr/archiveservice_generator/data"
	"github.com/etiennelndr/archiveservice_generator/utils"
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

// InitDirectories create directories and files for each service
func (g *Generator) InitDirectories() error {
	path := "../../../../../"
	testpath := path + "Tests/"
	filepath, err := filepath.Abs(testpath)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath, os.ModePerm)
	if err != nil {
		return err
	}

	for _, service := range g.GenArea.Services {
		// nameservice
		serviceabspath := testpath + strings.ToLower(service.Name) + "service/"
		err = os.MkdirAll(serviceabspath, os.ModePerm)
		if err != nil {
			return err
		}

		// name
		name := serviceabspath + strings.ToLower(service.Name) + "/"
		err = os.MkdirAll(name, os.ModePerm)
		if err != nil {
			return err
		}

		// service
		err = os.MkdirAll(name+"service/", os.ModePerm)
		if err != nil {
			return err
		}
		f, err := os.Create(name + "service/service.go")
		if err != nil {
			return err
		}
		utils.WriteHeader(f, "service")
		f.Close()

		// consumer
		err = os.MkdirAll(name+"consumer/", os.ModePerm)
		if err != nil {
			return err
		}
		f, err = os.Create(name + "consumer/consumer.go")
		if err != nil {
			return err
		}
		utils.WriteHeader(f, "consumer")
		f.Close()

		// provider
		err = os.MkdirAll(name+"provider/", os.ModePerm)
		if err != nil {
			return err
		}
		f, err = os.Create(name + "provider/provider.go")
		if err != nil {
			return err
		}
		utils.WriteHeader(f, "provider")
		f.Close()

		// data
		err = os.MkdirAll(serviceabspath+"data/", os.ModePerm)
		if err != nil {
			return err
		}
		f, err = os.Create(serviceabspath + "data/data.go")
		if err != nil {
			return err
		}
		utils.WriteHeader(f, "data")
		f.Close()

		// errors
		err = os.MkdirAll(serviceabspath+"errors/", os.ModePerm)
		if err != nil {
			return err
		}
		f, err = os.Create(serviceabspath + "errors/errors.go")
		if err != nil {
			return err
		}
		utils.WriteHeader(f, "errors")
		f.Close()

		// tests
		err = os.MkdirAll(serviceabspath+"tests/", os.ModePerm)
		if err != nil {
			return err
		}
		f, err = os.Create(serviceabspath + "tests/tests.go")
		if err != nil {
			return err
		}
		utils.WriteHeader(f, "tests")
		f.Close()
	}

	return nil
}

// CreateService TODO:
func (g *Generator) CreateService() error {
	return nil
}

// CreateProvider TODO:
func (g *Generator) CreateProvider() error {
	return nil
}

// CreateConsumer TODO:
func (g *Generator) CreateConsumer() error {
	return nil
}

// CreateData TODO:
func (g *Generator) CreateData() error {
	return nil
}

// CreateErrors TODO:
func (g *Generator) CreateErrors() error {
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

			// Retrieve all of the operations
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

			// Retrieve the service data types
			for _, comp := range service.Datas.Composites {
				c := Composite{
					Name:               comp.Name,
					Comment:            comp.Comment,
					ShortFormPart:      comp.ShortFormPart,
					NameOfTypeToExtend: comp.Extend.TypeToExtend.Name,
					AreaOfTypeToExtend: comp.Extend.TypeToExtend.Area,
				}
				for _, field := range comp.Fields {
					f := Field{
						CanBeNull: field.FieldCanBeNull,
						Comment:   field.Comment,
						Name:      field.Name,
						TypeArea:  field.FieldType.Area,
						TypeName:  field.FieldType.Name,
					}
					c.AddField(f)
				}
				s.AddComposite(c)
			}

			for _, enum := range service.Datas.Enumerations {
				e := Enumeration{
					Comment:       enum.Comment,
					Name:          enum.Name,
					ShortFormPart: enum.ShortFormPart,
				}
				for _, item := range enum.Items {
					i := Item{
						Comment: item.Comment,
						NValue:  item.NValue,
						Value:   item.Value,
					}
					e.AddItem(i)
				}
				s.AddEnumeration(e)
			}

			// Store this service in the area
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
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		send.AddType(data)
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
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		submit.AddType(data)
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
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		request.AddType(data)
	}
	op.Pattern.AddMessage(request)

	// Response Message
	response := Message{
		Name: "response",
	}
	for _, t := range operation.Message.Response.Types {
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		response.AddType(data)
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
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		invoke.AddType(data)
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
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		response.AddType(data)
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
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		progress.AddType(data)
	}
	op.Pattern.AddMessage(progress)

	// Invoke Message
	update := Message{
		Name: "update",
	}
	for _, t := range operation.Message.Update.Types {
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		update.AddType(data)
	}
	op.Pattern.AddMessage(update)

	// Response Message
	response := Message{
		Name: "response",
	}
	for _, t := range operation.Message.Response.Types {
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		response.AddType(data)
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
		data := Type{
			Area:    t.Area,
			Name:    t.Name,
			List:    t.List,
			Service: t.Service,
		}
		publishNotify.AddType(data)
	}
	op.Pattern.AddMessage(publishNotify)

	// Add this new operation to the service
	s.AddOperation(op)
}
