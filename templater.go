package main

import (
    "bytes"
    "log"
    "strings"
    "text/template"
)

var (
	TypesTemplate *template.Template
)
func CreateTemplates() {
    TypesTemplate = template.Must(template.New("base_template").
        Funcs(template.FuncMap{ "ConvertSoapTypeToTypescript": ConvertSoapTypeToTypescript }).
        ParseFiles("templates/types.gohtml"))
}

func ConvertSoapTypeToTypescript(soapType string) string {
    if strings.Contains(soapType, ":") {
        soapType = strings.Split(soapType, ":")[1]
    }

    tsTypes := map[string]string{
        "base64Binary":  "string",
        "byte":          "number",
        "dateTime":      "string",
        "double":        "number",
        "float":         "number",
        "hexBinary":     "string",
        "int":           "number",
        "short":         "number",
        "signedInt":     "number",
        "unsignedByte":  "number",
        "unsignedInt":   "number",
        "unsignedShort": "number",
    }
    if value, present := tsTypes[soapType]; present {
        return value
    } else {
        return soapType
    }
}

func ExecuteTemplate(name string, data interface{}) string {
    var buffer bytes.Buffer
    err := TypesTemplate.ExecuteTemplate(&buffer, name, data)
    if err != nil {
        log.Panic(err)
    }

    return buffer.String()
}

func CreateInterfaceFromComplexType(complexType ComplexType) string {
    return ExecuteTemplate("complex_type_to_interface", complexType)
}
func CreateInterfaceFromElement(element Element) string {
    return ExecuteTemplate("element_to_interface", element)
}
