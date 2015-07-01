package doc

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	requestTmpl *template.Template
	requestFmt  = `{{if or .HasBody .HasHeader}}
+ Request {{if .HasContentType}}({{.Header.ContentType}}){{end}}{{with .Header}}
{{.Render}}{{end}}{{with .Body}}
{{.Render}}{{end}}
{{end}}`
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

func NewRequest(req *http.Request) (*Request, error) {
	body1, body2, err := cloneBody(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = nopCloser{body1}

	b2bytes := body2.Bytes()
	fmt.Println("BODY:", string(b2bytes))

	return &Request{
		Header: NewHeader(req.Header),
		Body:   b2bytes,
	}, nil
}

func (r *Request) Render() string {
	return render(requestTmpl, r)
}

func (r *Request) HasBody() bool {
	return len(r.Body) > 0
}

func (r *Request) HasHeader() bool {
	return r.Header != nil && len(r.Header.DisplayHeader) > 0
}

func (r *Request) HasContentType() bool {
	return r.Header != nil && len(r.Header.ContentType) > 0
}
