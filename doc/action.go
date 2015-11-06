package doc

import (
	"text/template"

	"github.com/adams-sarah/test2doc/doc/parse"
)

var (
	actionTmpl *template.Template
	actionFmt  = `### {{.Title}} [{{.Method}}]
{{.Description}}{{with .Request}}{{.Render}}{{end}}
{{with .Response}}{{.Render}}{{end}}`
)

func init() {
	actionTmpl = template.Must(template.New("action").Parse(actionFmt))
}

type Action struct {
	Title       string
	Description string
	Method      HTTPMethod
	Request     Request // status OK
	Response    Response

	// TODO: document non-OK requests ??
}

func (a *Action) Render() string {
	return render(actionTmpl, a)
}

func NewAction(req *Request, resp *Response, handlerName string) (*Action, error) {
	title := parse.GetTitle(handlerName)
	desc := parse.GetDescription(handlerName)

	return &Action{
		Title:       title,
		Description: desc,
		Method:      HTTPMethod(req.Method),
		Request:     *req,
		Response:    *resp,
	}, nil

}
