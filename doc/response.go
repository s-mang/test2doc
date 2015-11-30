package doc

import (
	"net/http/httptest"
	"text/template"
)

var (
	responseTmpl *template.Template
	responseFmt  = `+ Response {{.StatusCode}} {{if .HasContentType}}({{.Header.ContentType}}){{end}}{{with .Header}}

{{.Render}}{{end}}{{with .Body}}
{{.Render}}{{end}}
`
)

func init() {
	responseTmpl = template.Must(template.New("response").Parse(responseFmt))
}

type Response struct {
	StatusCode  int
	Description string
	Header      *Header
	Body        *Body

	// TODO:
	// Attributes
	// Schema
}

func NewResponse(resp *httptest.ResponseRecorder) *Response {
	content := resp.Body.Bytes()
	contentType := resp.Header().Get("Content-Type")

	return &Response{
		StatusCode: resp.Code,
		Header:     NewHeader(resp.Header()),
		Body:       NewBody(content, contentType),
	}
}

func (r *Response) Render() string {
	return render(responseTmpl, r)
}

func (r *Response) HasContentType() bool {
	return r.Header != nil && len(r.Header.ContentType) > 0
}
