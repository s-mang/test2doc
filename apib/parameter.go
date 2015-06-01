//go:generate stringer -type=ParameterType

package apib

import (
	"fmt"
	"text/template"
)

type ParameterType int

const (
	Number ParameterType = iota
	String
	Boolean
)

var (
	parameterTmpl *template.Template
	parameterFmt  = `	+ {{.Title}} {{.Value.Quote}} ({{.Type.String}}) - {{.Description}}`
)

func init() {
	parameterTmpl = template.Must(template.New("parameter").Parse(parameterFmt))
}

type Parameter struct {
	Title       string
	Description string
	Value       ParameterValue
	Type        ParameterType
	IsRequired  bool

	// TODO:
	// DefaultValue
}

func (p *Parameter) Render() string {
	return render(parameterTmpl, p)
}

type ParameterValue string

func (v ParameterValue) Quote() string {
	return fmt.Sprintf("`%s`", string(v))
}
