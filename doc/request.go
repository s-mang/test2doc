package doc

import (
	"net/http"
	"text/template"
)

var (
	requestTmpl *template.Template
	requestFmt  = `{{if or .Body .Header}}
+ Request {{with .Header}}({{.ContentType}})
{{.Render}}{{end}}
{{with .Body}}{{.Render}}{{end}}{{end}}`
)

func init() {
	requestTmpl = template.Must(template.New("request").Parse(requestFmt))
}

type Request struct {
	Header *Header
	Body   Body

	// TODO:
	// Attributes
	// Schema
}

func (r *Request) Render() string {
	return render(requestTmpl, r)
}

func NewRequest(req *http.Request) (*Request, error) {
	body1, body2, err := cloneBody(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = nopCloser{body1}

	return &Request{
		Header: NewHeader(req.Header),
		Body:   body2.Bytes(),
	}, nil
}
