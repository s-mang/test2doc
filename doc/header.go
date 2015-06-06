package doc

import (
	"net/http"
	"text/template"
)

var (
	headerTmpl *template.Template
	headerFmt  = `
	+ Headers
		{{range $name, $vals := .}}
		{{$name}}: {{$vals | commaJoin}}{{end}}
`
)

func init() {
	funcMap := template.FuncMap{
		"commaJoin": commaJoin,
	}

	headerTmpl = template.Must(template.New("headers").Funcs(funcMap).Parse(headerFmt))
}

type Header http.Header

func (h Header) Render() string {
	return render(headerTmpl, h)
}

func (h Header) ContentType() string {
	return http.Header(h).Get("Content-Type")
}
