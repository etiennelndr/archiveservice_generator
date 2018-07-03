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
	filepath, err := testPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath, os.ModePerm)
	if err != nil {
		return err
	}

	for _, service := range g.GenArea.Services {
		// nameservice
		serviceabspath := filepath + "/" + strings.ToLower(service.Name) + "service/"
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

		// constants
		err = os.MkdirAll(name+"constants/", os.ModePerm)
		if err != nil {
			return err
		}
		f, err = os.Create(name + "constants/constants.go")
		if err != nil {
			return err
		}
		utils.WriteHeader(f, "constants")
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

// CreateInformation TODO:
func (g *Generator) CreateInformation() error {
	err := g.createConstants()

	err = g.createService()
	if err != nil {
		return err
	}

	err = g.createData()
	if err != nil {
		return err
	}

	err = g.createErrors()
	if err != nil {
		return err
	}

	err = g.createProvider()
	if err != nil {
		return err
	}

	return g.createConsumer()
}

func (g *Generator) createConstants() error {
	filepath, err := testPath()
	if err != nil {
		return err
	}

	for _, s := range g.GenArea.Services {
		var buffer = new(bytes.Buffer)
		serviceNameToLower := strings.ToLower(s.Name)
		constantsfile := filepath + "/" + serviceNameToLower + "service/" + serviceNameToLower + "/constants/constants.go"

		file, err := os.OpenFile(constantsfile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}
		defer file.Close()

		buffer.WriteString("\n// Constants for the " + s.Name + " Service\n")
		buffer.WriteString("const (\n")
		buffer.WriteString("\t" + serviceIdentifier(s) + " = \"" + s.Name + "\"\n")
		buffer.WriteString("\t" + serviceNumber(s) + "     = " + s.Number + "\n")
		buffer.WriteString(")\n")
		buffer.WriteString("\nconst (\n")
		buffer.WriteString("\t" + areaIdentifier(s) + " = \"" + g.GenArea.Name + "\"\n")
		buffer.WriteString(")\n")
		if len(s.Operations) != 0 {
			buffer.WriteString("\n// Constants for the operations\n")
			buffer.WriteString("const (\n")
			for _, op := range s.Operations {
				buffer.WriteString("\tOPERATION_IDENTIFIER_" + strings.ToUpper(op.Name) + " = " + op.Number + "\n")
			}
			buffer.WriteString(")\n")
		}

		_, err = file.Write(buffer.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}

func serviceImports(buf *bytes.Buffer, s Service) {
	sName := strings.ToLower(s.Name)
	buf.WriteString("\nimport (\n")
	buf.WriteString("\t\"github.com/ccsdsmo/malgo/mal\"\n") // FIXME: it might not be generic
	buf.WriteString("\t\"github.com/ccsdsmo/malgo/com\"\n") // FIXME: same as mal import
	buf.WriteString("\tcnst \"github.com/etiennelndr/tests/" + sName + "service/" + sName + "/constants\"")
	buf.WriteString("\n\t\"sync\"\n")
	buf.WriteString(")\n")
}

func serviceStructure(buf *bytes.Buffer, serviceName string) {
	buf.WriteString("\ntype " + serviceName + "Service struct {\n")
	buf.WriteString("\tAreaIdentifier 	 mal.Identifier\n")
	buf.WriteString("\tServiceIdentifier mal.Identifier\n")
	buf.WriteString("\tAreaNumber 		 mal.UShort\n")
	buf.WriteString("\tServiceNumber 	 mal.Integer\n")
	buf.WriteString("\tAreaVersion 		 mal.UOctet\n\n")
	buf.WriteString("\trunning 			 bool\n")
	buf.WriteString("\twg 				 sync.WaitGroup\n")
	buf.WriteString("}\n")
}

func serviceCreateService(buf *bytes.Buffer, s Service, area Area) {
	buf.WriteString("\nfunc New" + s.Name + "Service() *" + s.Name + "Service {\n")
	buf.WriteString("\t" + strings.ToLower(s.Name) + "Service := &" + s.Name + "Service{\n")
	buf.WriteString("\t\tAreaIdentifier: cnst." + areaIdentifier(s) + ",\n")
	buf.WriteString("\t\tServiceIdentifier: cnst." + serviceIdentifier(s) + ",\n")
	buf.WriteString("\t\tAreaNumber: com." + strings.ToUpper(area.Name) + "_AREA_NUMBER,\n")
	buf.WriteString("\t\tServiceNumber: cnst." + serviceNumber(s) + ",\n")
	buf.WriteString("\t\tAreaVersion: com." + strings.ToUpper(area.Name) + "_AREA_VERSION,\n")
	buf.WriteString("\t\trunning: true,\n")
	buf.WriteString("\t\twg: *new(sync.WaitGroup),\n")
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn " + strings.ToLower(s.Name) + "Service\n")
	buf.WriteString("}\n")
}

func serviceOperations(buf *bytes.Buffer, s Service) {
	for _, op := range s.Operations {
		buf.WriteString("\n")
		// Print the comment of the operation
		printComment(buf, op.Name+": "+op.Comment)

		// Now print the header and some lines
		buf.WriteString("func (s *" + s.Name + "Service) " + charsToUpper(op.Name, 0) + " (consumerURL string, providerURL string,")
		for i, t := range op.Pattern.Messages[0].Types {
			//println
			buf.WriteString(" " + charsToLower(t.AdaptType(), 0) + " " + t.AdaptType())
			if i+1 < len(op.Pattern.Messages[0].Types) {
				buf.WriteString(",")
			}
		}
		buf.WriteString(") (")
		for _, t := range op.Pattern.Messages[len(op.Pattern.Messages)-1].Types {
			if t.Name == "Element" {
				// TODO: finish this section
			}
			buf.WriteString("*" + strings.ToLower(t.Area) + "." + charsToUpper(t.Name, 0) + ", ")
		}
		buf.WriteString("error) {\n")
		// Elements to return
		buf.WriteString("\treturn nil\n")
		buf.WriteString("}\n")
	}
}

func printComment(buf *bytes.Buffer, comment string) {
	comment = strings.Replace(comment, "\n", " ", -1)
	splitComment := strings.Split(comment, " ")
	splitComment = splitComment[:len(splitComment)-1]

	// Uppercase the first letter
	splitComment[0] = charsToUpper(splitComment[0], 0)
	// Put a space between the function name and the ':'
	splitComment[0] = strings.Replace(splitComment[0], ":", " :", -1)

	var stop = false
	for i := 0; i < len(splitComment); i++ {
		var str string
		stop = false
		for !stop {
			if len(str+" "+splitComment[i]) < 64 {
				str += " " + splitComment[i]
				if i+1 < len(splitComment) {
					i++
				} else {
					break
				}
			} else {
				stop = true
			}
		}
		buf.WriteString("//" + str + "\n")
	}
	if stop {
		buf.WriteString("// " + splitComment[len(splitComment)-1] + "\n")
	}
}

func (g *Generator) createService() error {
	filepath, err := testPath()
	if err != nil {
		return err
	}

	for _, service := range g.GenArea.Services {
		var buffer = new(bytes.Buffer)
		serviceNameToLower := strings.ToLower(service.Name)
		servicefile := filepath + "/" + serviceNameToLower + "service/" + serviceNameToLower + "/service/service.go"

		file, err := os.OpenFile(servicefile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}
		defer file.Close()

		// TODO: Create imports
		serviceImports(buffer, service)

		// Create the structure for the Service
		serviceStructure(buffer, service.Name)

		// A method to create a new service
		serviceCreateService(buffer, service, g.GenArea)

		// Create the operations for each service
		serviceOperations(buffer, service)

		_, err = file.Write(buffer.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) createProvider() error {
	for _, service := range g.GenArea.Services {
		fmt.Println("> Provider: " + service.Name)
	}
	return nil
}

func (g *Generator) createConsumer() error {
	for _, service := range g.GenArea.Services {
		fmt.Println("> Consumer: " + service.Name)
	}
	return nil
}

func (g *Generator) createData() error {
	for _, data := range g.GenArea.Composites {
		fmt.Println("> Data: " + data.Name)
	}
	return nil
}

func (g *Generator) createErrors() error {
	for _, err := range g.GenArea.Errors {
		fmt.Println("> Error: " + err.Name)
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
			comp := createComposite(composite)
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
				c := createComposite(comp)
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

func createComposite(composite data.Composite) Composite {
	c := Composite{
		Name:               composite.Name,
		Comment:            composite.Comment,
		ShortFormPart:      composite.ShortFormPart,
		NameOfTypeToExtend: composite.Extend.TypeToExtend.Name,
		AreaOfTypeToExtend: composite.Extend.TypeToExtend.Area,
	}
	for _, field := range composite.Fields {
		f := Field{
			CanBeNull: field.FieldCanBeNull,
			Comment:   field.Comment,
			Name:      field.Name,
			TypeArea:  field.FieldType.Area,
			TypeName:  field.FieldType.Name,
		}
		c.AddField(f)
	}

	return c
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

func testPath() (string, error) {
	path := "../"
	testpath := path + "tests/"
	filepath, err := filepath.Abs(testpath)
	if err != nil {
		return "", err
	}
	return filepath, nil
}

func serviceIdentifier(s Service) string {
	return strings.ToUpper(s.Name) + "_SERVICE_SERVICE_IDENTIFIER"
}

func serviceNumber(s Service) string {
	return strings.ToUpper(s.Name) + "_SERVICE_SERVICE_NUMBER"
}

func areaIdentifier(s Service) string {
	return strings.ToUpper(s.Name) + "_SERVICE_AREA_IDENTIFIER"
}

func charsToLower(str string, pos ...int) string {
	splitstr := strings.Split(str, "")
	for i := 0; i < len(pos); i++ {
		splitstr[i] = strings.ToLower(splitstr[i])
	}

	return strings.Join(splitstr, "")
}

func charsToUpper(str string, pos ...int) string {
	splitstr := strings.Split(str, "")
	for i := 0; i < len(pos); i++ {
		splitstr[i] = strings.ToUpper(splitstr[i])
	}

	return strings.Join(splitstr, "")
}
