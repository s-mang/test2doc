package doc

import "text/template"

var (
	actionTmpl *template.Template
	actionFmt  = `### {{.Title}} [{{.Method}}]
{{.Description}}
{{with .Request}}{{.Render}}{{end}}
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

func NewAction(method HTTPMethod, req *Request, resp *Response) (*Action, error) {
	return &Action{
		Title:       "Some Action",
		Description: "Some description.",
		Method:      method,
		Request:     *req,
		Response:    *resp,
	}, nil

}
