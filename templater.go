package main

import (
    "bytes"
    "log"
    "strings"
    "text/template"
)

var (
	BaseTemplate *template.Template
)
func CreateBaseTemplate() {
    BaseTemplate = template.Must(template.New("base_template").
        Funcs(template.FuncMap{ "ConvertSoapTypeToTypescript": ConvertSoapTypeToTypescript }).
        ParseFiles("template.gohtml"))
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
    err := BaseTemplate.ExecuteTemplate(&buffer, name, data)
    if err != nil {
        log.Panic(err)
    }

    return buffer.String()
}

func CreateInterfaceFromComplexType(complexType ComplexType) string {
    return ExecuteTemplate("complex_type_to_interface", complexType)
}
