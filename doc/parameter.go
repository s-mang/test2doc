package doc

import (
	"fmt"
	"regexp"
	"text/template"
	"strings"
	"reflect"

	"github.com/everytv/test2doc/doc/parse"
)

type ParameterType int

const (
	Number ParameterType = iota
	String
	Boolean
)

const (
	numberRe = `^[0-9\.]+$`
	boolRe   = `^(?:[tT][rR][uU][eE]|[fF][aA][lL][sS][eE])$`
)

var (
	parameterTmpl *template.Template
	parameterFmt  = "    + {{.Name}}: {{.Value.Quote}} ({{.Type.String}}, {{with .IsRequired}}required{{else}}optional{{end}}){{with .Description}} - {{.}}{{end}}{{with .DefaultValue}}\n      + Default: {{.}}{{end}}"
)

func init() {
	parameterTmpl = template.Must(template.New("parameter").Parse(parameterFmt))
}

type Parameter struct {
	Name        string
	Description string
	Value       ParameterValue
	Type        ParameterType
	IsRequired  bool

	DefaultValue string
}

func MakeParameter(key, val string) Parameter {
	description, isRequired, defaultValue := getPropertyOf(key)
	return Parameter{
		Name:       key,
		Description: description,
		Value:      ParameterValue(val),
		Type:       paramType(val),
		IsRequired: isRequired, // assume anything in route URL is required
		// query params are a different story
		DefaultValue: defaultValue,
	}
}

func (p *Parameter) Render() string {
	return render(parameterTmpl, p)
}

type ParameterValue string

func (val ParameterValue) Quote() (qval string) {
	if len(val) > 0 {
		qval = fmt.Sprintf("`%s`", string(val))
	}

	return
}

func paramType(val string) ParameterType {
	if isBool(val) {
		return Boolean
	} else if isNumber(val) {
		return Number
	} else {
		return String
	}
}

func isBool(str string) bool {
	re := regexp.MustCompile(boolRe)
	return re.MatchString(str)
}

func isNumber(str string) bool {
	re := regexp.MustCompile(numberRe)
	return re.MatchString(str)
}

func getPropertyOf(key string) (description string, isRequired bool, defaultValue string) {
	if parse.ParamsType == nil {
		return
	}

	paramsType := *parse.ParamsType
	var field *reflect.StructField
	for i := 0; i < paramsType.NumField(); i++ {
		jsonTag := paramsType.Field(i).Tag.Get("json")
		if jsonTag == key {
			wk := paramsType.Field(i)
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
		kv := strings.Split(pair, "=")
		if kv[0] == "description" && len(kv) >= 2 {
			description = kv[1]
		} else if kv[0] == "required" {
			isRequired = true
		} else if kv[0] == "default" && len(kv) >= 2 {
			defaultValue = kv[1]
		} else {
			// shoddy validation
			panic(fmt.Sprintf("unknown format: %v", pair))
		}
	}

	return
}

func (pt ParameterType) String() string {
	switch pt {
	case Number:
		return "number"
	case Boolean:
		return "boolean"
	default:
		return "string"
	}
}
