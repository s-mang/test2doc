package doc

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/adams-sarah/test2doc/doc/parse"
)

var (
	actionTmpl *template.Template
	actionFmt  = `### {{.Title}} [{{.Method}}]
{{.Description}}{{range $req := .Requests}}
{{with $req}}{{.Render}}
{{.Response.Render}}{{end}}{{end}}`
)

func init() {
	actionTmpl = template.Must(template.New("action").Parse(actionFmt))
}

type Action struct {
	Title       string
	Description string
	Method      HTTPMethod
	Requests    []*Request
}

func (a *Action) Render() string {
	reqsMap := map[int][]*Request{}
	for i, req := range a.Requests {
		if reqsMap[req.Response.StatusCode] == nil {
			reqsMap[req.Response.StatusCode] = []*Request{}
		}

		reqsMap[req.Response.StatusCode] = append(reqsMap[req.Response.StatusCode], a.Requests[i])
	}

	sortedReqs := reqsMap[http.StatusOK]
	delete(reqsMap, http.StatusOK)

	for _, reqs := range reqsMap {
		sortedReqs = append(sortedReqs, reqs...)
	}

	a.Requests = sortedReqs

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
		Requests:    []*Request{},
	}, nil

}

func (a *Action) AddRequest(req *Request, resp *Response) {
	if req.Description != "" {
		a.Description = req.Description
	}
	if req.Title != "" {
		a.Title = req.Title
	}
	req.Response = resp
	a.Requests = append(a.Requests, req)
}
