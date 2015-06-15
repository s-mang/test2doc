package doc

import (
	"net/http/httptest"
	"text/template"
)

var (
	responseTmpl *template.Template
	responseFmt  = `+ Response {{.StatusCode}} {{with .Header}}({{.ContentType}})
{{.Render}}{{end}}{{with .Body}}
{{.Render}}{{end}}`
)

func init() {
	responseTmpl = template.Must(template.New("response").Parse(responseFmt))
}

type Response struct {
	StatusCode  int
	Description string
	Header      *Header
	Body        Body

	// TODO:
	// Attributes
	// Schema
}

func NewResponse(resp *httptest.ResponseRecorder) *Response {
	return &Response{
		StatusCode: resp.Code,
		Header:     NewHeader(resp.Header()),
		Body:       resp.Body.Bytes(),
	}
}

func (r *Response) Render() string {
	return render(responseTmpl, r)
}
