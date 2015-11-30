package doc

import (
	"strings"
	"text/template"

	"github.com/adams-sarah/test2doc/doc/parse"
)

var (
	actionTmpl *template.Template
	actionFmt  = `### {{.Title}} [{{.Method}}]
{{.Description}}{{range $req, $resp := .Requests}}
{{with $req}}{{.Render}}{{end}}
{{with $resp}}{{.Render}}{{end}}{{end}}`
)

func init() {
	actionTmpl = template.Must(template.New("action").Parse(actionFmt))
}

type Action struct {
	Title       string
	Description string
	Method      HTTPMethod
	Requests    map[*Request]*Response

	// TODO: document non-OK requests ??
}

func (a *Action) Render() string {
	return render(actionTmpl, a)
}

func NewAction(method, handlerName string) (*Action, error) {
	title := parse.GetTitle(handlerName)
	if len(title) == 0 {
		title = strings.Title(method)
	}

	desc := parse.GetDescription(handlerName)

	return &Action{
		Title:       title,
		Description: desc,
		Method:      HTTPMethod(method),
		Requests:    map[*Request]*Response{},
	}, nil

}

func (a *Action) AddRequest(req *Request, resp *Response) {
	a.Requests[req] = resp
}
