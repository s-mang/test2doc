package doc

import (
	"net/http"
	"net/http/httptest"
	"text/template"
)

var (
	actionTmpl *template.Template
	actionFmt  = `### {{.Title}} [{{.Method}}]
{{.Description}}{{with .Request}}
{{.Render}}{{end}}{{with .Response}}
{{.Render}}{{end}}`
)

func init() {
	actionTmpl = template.Must(template.New("action").Parse(actionFmt))
}

type Action struct {
	Title       string
	Description string
	Method      httpMethod
	Request     Request // status OK
	Response    Response

	// TODO: document non-OK requests ??
}

func (a *Action) Render() string {
	return render(actionTmpl, a)
}

func NewAction(req *http.Request, resp *httptest.ResponseRecorder) (*Action, error) {
	docReq, err := NewRequest(req)
	if err != nil {
		return nil, err
	}

	docResp := NewResponse(resp)

	return &Action{
		Title:       "Some Action",
		Description: "Some description.",
		Method:      httpMethod(req.Method),
		Request:     *docReq,
		Response:    *docResp,
	}, nil

}
