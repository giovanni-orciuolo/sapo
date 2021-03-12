package main

import "encoding/xml"

type Definitions struct {
	XMLName     xml.Name      `xml:"definitions"`
	Name        string        `xml:"name,attr"`
	Types       Types         `xml:"types"`
	Messages    []*Message    `xml:"message"`
	PortTypes   []*PortType   `xml:"portType"`
	Bindings    []*Binding    `xml:"binding"`
	Services    []*Service    `xml:"service"`
	SimpleTypes []*SimpleType `xml:"simpleType"`
}

type Types struct {
	XMLName xml.Name  `xml:"types"`
	Schemas []*Schema `xml:"schema"`
}
type Schema struct {
	XMLName            xml.Name       `xml:"schema"`
	ElementFormDefault string         `xml:"elementFormDefault,attr"`
	TargetNamespace    string         `xml:"targetNamespace,attr"`
	Imports            []*Import      `xml:"import"`
	ComplexTypes       []*ComplexType `xml:"complexType"`
	Elements           []*Element     `xml:"element"`
}
type Import struct {
	XMLName   xml.Name `xml:"import"`
	Namespace string   `xml:"namespace,attr"`
}
type Element struct {
	XMLName     xml.Name     `xml:"element"`
	Name        string       `xml:"name,attr"`
	Type        string       `xml:"type,attr"`
	MinOccurs   string       `xml:"minOccurs,attr"`
	MaxOccurs   string       `xml:"maxOccurs,attr"`
	Nillable    bool         `xml:"nillable,attr"`
	ComplexType *ComplexType `xml:"complexType"`
}
type ComplexType struct {
	XMLName    xml.Name     `xml:"complexType"`
	Name       string       `xml:"name,attr"`
	Sequence   *Sequence    `xml:"sequence"`
	Attributes []*Attribute `xml:"attribute"`
}
type Sequence struct {
	XMLName  xml.Name   `xml:"sequence"`
	Elements []*Element `xml:"element"`
}
type Attribute struct {
	XMLName    xml.Name    `xml:"attribute"`
	Name       string      `xml:"name,attr"`
	Type       string      `xml:"type,attr"`
	Use        string      `xml:"use,attr"`
	SimpleType *SimpleType `xml:"simpleType"`
}
type SimpleType struct {
	XMLName     xml.Name     `xml:"simpleType"`
	Name        string       `xml:"name"`
	Restriction *Restriction `xml:"restriction"`
}
type Restriction struct {
	XMLName      xml.Name     `xml:"restriction"`
	Base         string       `xml:"base"`
	Length       int          `xml:"length"`
	Pattern      string       `xml:"pattern"`
	MinLength    int          `xml:"minLength"`
	MaxLength    int          `xml:"maxLength"`
	Enumerations *Enumeration `xml:"enumeration"`
}
type Enumeration struct {
	XMLName    xml.Name   `xml:"enumeration"`
	Value      string     `xml:"value"`
	Annotation Annotation `xml:"annotation"`
}
type Annotation struct {
	Documentation string `xml:"documentation"`
}

type Message struct {
	XMLName xml.Name `xml:"message"`
	Name    string   `xml:"name,attr"`
	Parts   []*Part  `xml:"part"`
}
type Part struct {
	XMLName xml.Name `xml:"part"`
	Name    string   `xml:"name,attr"`
	Element string   `xml:"element,attr"`
}

type PortType struct {
	XMLName    xml.Name     `xml:"portType"`
	Name       string       `xml:"name,attr"`
	Operations []*Operation `xml:"operation"`
}
type Operation struct {
	XMLName       xml.Name       `xml:"operation"`
	Name          string         `xml:"name,attr"`
	Input         *Input         `xml:"input"`
	Output        *Output        `xml:"output"`
	SoapOperation *SoapOperation `xml:"operation"`
}
type Input struct {
	XMLName xml.Name `xml:"input"`
	Action  string   `xml:"Action,attr"`
	Message string   `xml:"message,attr"`
	Body    *Body    `xml:"body"`
}
type Output struct {
	XMLName xml.Name `xml:"output"`
	Action  string   `xml:"Action,attr"`
	Message string   `xml:"message,attr"`
}
type SoapOperation struct {
	XMLName    xml.Name `xml:"operation"`
	SoapAction string   `xml:"soapAction,attr"`
	Style      string   `xml:"style,attr"`
}
type Body struct {
	XMLName xml.Name `xml:"body"`
	Use     string   `xml:"use,attr"`
}

type Binding struct {
	XMLName     xml.Name     `xml:"binding"`
	Name        string       `xml:"name,attr"`
	Type        string       `xml:"type,attr"`
	SoapBinding *SoapBinding `xml:"binding"`
	Operations  []*Operation `xml:"operation"`
}
type SoapBinding struct {
	XMLName   xml.Name `xml:"binding"`
	Transport string   `xml:"transport,attr"`
}

type Service struct {
	XMLName xml.Name `xml:"service"`
	Name    string   `xml:"name,attr"`
	Port    *Port    `xml:"port"`
}
type Port struct {
	XMLName xml.Name `xml:"port"`
	Name    string   `xml:"name,attr"`
	Binding string   `xml:"binding,attr"`
	Address *Address `xml:"address"`
}
type Address struct {
	XMLName  xml.Name `xml:"address"`
	Location string   `xml:"location,attr"`
}
