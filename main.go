package main

import (
    "encoding/xml"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    wsdlPath := os.Args[1]
    if wsdlPath == "" {
        log.Fatal("Please specify a WSDL path or URL.")
        return
    }

    xmlFile, err := os.Open(wsdlPath)
    if err != nil {
        log.Fatal(err)
        return
    }

    defer xmlFile.Close()
    bytes, _ := ioutil.ReadAll(xmlFile)

    var definitions Definitions
    err = xml.Unmarshal(bytes, &definitions)
    if err != nil {
        log.Fatal(err)
        return
    }

    log.Println("Parsed WSDL!")
    CreateBaseTemplate()

    for _, complexType := range definitions.Types.Schemas[0].ComplexTypes {
        iface := CreateInterfaceFromComplexType(*complexType)
        log.Println(iface)
    }
}
