package doc

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/everytv/test2doc/doc/parse"
)

var (
	dataStructuresTmpl *template.Template
	dataStructuresFmt  = `
# Data Structures

## {{.Name}} (object){{if .Attributes}}
+ data
    + {{.ObjectName}} ({{.ObjectName}})
+ Include Meta

## {{.ObjectName}} (object)
{{range .Attributes}}{{.Render}}
{{end}}{{end}}
`
)

func init() {
	dataStructuresTmpl = template.Must(template.New("datastructures").Parse(dataStructuresFmt))
}

type DataStructure struct {
	Name       string
	ObjectName string
	Attributes []Attribute
}

func (d *DataStructure) Render() string {
	return render(dataStructuresTmpl, d)
}

func getPropertyOfResponse(key string) (description string, isRequired bool, defaultValue string) {
	if parse.ResponseType == nil {
		return
	}

	responseType := *parse.ResponseType
	var field *reflect.StructField
	for i := 0; i < responseType.NumField(); i++ {
		jsonTag := responseType.Field(i).Tag.Get("json")
		if jsonTag == key {
			wk := responseType.Field(i)
			field = &wk
			break
		}
	}
	if field == nil {
		return
	}

	// `apidoc:"required,default=...,description=..."`
	apidocTag := field.Tag.Get("apidoc")
	for _, pair := range strings.Split(apidocTag, ",") {
		if len(pair) == 0 {
			continue
		}
		kv := strings.Split(pair, "=")
		if kv[0] == "description" && len(kv) == 2 {
			description = kv[1]
		} else if kv[0] == "required" && len(kv) == 1 {
			isRequired = true
		} else if kv[0] == "default" && len(kv) == 2 {
			defaultValue = kv[1]
		} else {
			// shoddy validation
			panic(fmt.Sprintf("unknown format: %v", pair))
		}
	}

	return
}
