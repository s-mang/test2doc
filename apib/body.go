package apib

import "text/template"

var (
	bodyTmpl *template.Template
	bodyFmt  = `
	+ Body

        {{.BodyStr}}
        
`
)

func init() {
	bodyTmpl = template.Must(template.New("body").Parse(bodyFmt))
}

type Body []byte

func (b *Body) Render() string {
	return render(bodyTmpl, b)
}
