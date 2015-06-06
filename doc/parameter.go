//go:generate stringer -type=ParameterType

package doc

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
	parameterFmt  = `	+ {{.Name}} {{.Value.Quote}} ({{.Type.String}}) - {{.Description}}`
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
