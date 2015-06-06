package apib

import "text/template"

var (
	actionTmpl *template.Template
	actionFmt  = `### {{.Title}} [{{.HTTPMethod}}]
{{.Description}}
{{with .Request}}
{{.Render}}{{end}}
{{with .Response}}
{{.Render}}{{end}}`
)

func init() {
	actionTmpl = template.Must(template.New("action").Parse(actionFmt))
}

type Action struct {
	Title       string
	Description string
	HTTPMethod  string
	Request     *Request // status OK
	Response    *Response

	// TODO: document non-OK requests ??
}

func (a *Action) Render() string {
	return render(actionTmpl, a)
}
