package doc

import (
	"text/template"
)

var (
	attributeTmpl *template.Template
	attributeFmt  = `
+ {{.Name}}{{with .IsObject}}{{else}}:{{end}} {{.Value.Quote}} ({{with .IsObject}}hogeType{{else}}{{.Type.String}}{{end}}, {{with .IsRequired}}required{{else}}optional{{end}}){{with .Description}} - {{.}}{{end}}{{with .DefaultValue}}
+ Default: {{.}}{{end}}
`
)

func init() {
	attributeTmpl = template.Must(template.New("attribute").Parse(attributeFmt))
}

type Attribute Parameter

func (p *Attribute) Render() string {
	return render(attributeTmpl, p)
}
