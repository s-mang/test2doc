package apib

import (
	"net/http"
	"text/template"
)

var (
	requestTmpl *template.Template
	requestFmt  = `{{if or .Header .Body}}
+ Request {{with .Header.ContentType}}({{.}}){{end}}
{{end}}
{{.Header.Render}}
{{.Body.Render}}
`
)

func init() {
	requestTmpl = template.Must(template.New("request").Parse(requestFmt))
}

type Request struct {
	Header Header
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
		Header: Header(req.Header),
		Body:   body2.Bytes(),
	}, nil
}

func (r *Request) BodyStr() (body string) {
	body = string(r.Body)
	if r.Header.ContentType() == "application/json" {
		var err error
		body, err = indentJSONBody(string(r.Body))
		if err != nil {
			panic(err.Error())
		}
	}

	return

}
