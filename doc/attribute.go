package doc

import (
	"text/template"
)

var (
	attributeTmpl *template.Template
	attributeFmt  = "        + {{.Name}}: {{.Value.Quote}} ({{.Type.String}}, {{with .IsRequired}}required{{else}}optional{{end}}){{with .Description}} - {{.}}{{end}}{{with .DefaultValue}}\n            + Default: {{.}}{{end}}"
)

func init() {
	attributeTmpl = template.Must(template.New("attribute").Parse(attributeFmt))
}

type Attribute Parameter

func (p *Attribute) Render() string {
	return render(attributeTmpl, p)
}
