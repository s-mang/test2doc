package doc

import (
	"fmt"
	"regexp"
	"text/template"
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
	parameterFmt  = `    + {{.Name}}: {{.Value.Quote}} ({{.Type.String}}){{with .Description}} - {{.}}{{end}}`
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

	// TODO:
	// DefaultValue
}

func MakeParameter(key, val string) Parameter {
	return Parameter{
		Name:       key,
		Value:      ParameterValue(val),
		Type:       paramType(val),
		IsRequired: true, // assume anything in route URL is required
		// query params are a different story
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
