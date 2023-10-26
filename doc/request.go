package doc

import (
	"net/http"
	"text/template"
)

var (
	requestTmpl *template.Template
	requestFmt  = `{{if or .HasBody .HasHeader}}
+ Request {{if .HasContentType}}({{.Header.ContentType}}){{end}}{{with .Header}}

{{.Render}}{{end}}{{with .Body}}
{{.Render}}{{end}}{{end}}`
)

func init() {
	requestTmpl = template.Must(template.New("request").Parse(requestFmt))
}

const (
	descriptionHeader = "X-Test2Doc-Description"
	titleHeader       = "X-Test2Doc-Title"
)

type Request struct {
	Header   *Header
	Body     *Body
	Method   string
	Response *Response

	// Headers which are pulled out of request
	// to use for generating documentation.
	Description string
	Title       string

	// TODO:
	// Attributes
	// Schema
}

func NewRequest(req *http.Request) (*Request, error) {
	// pull test2doc headers out of request,
	// then delete header so that it's not used
	// in the actual request
	desc := req.Header.Get(descriptionHeader)
	req.Header.Del(descriptionHeader)
	title := req.Header.Get(titleHeader)
	req.Header.Del(titleHeader)

	body1, body2, err := cloneBody(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = nopCloser{body1}

	b2bytes := body2.Bytes()
	contentType := req.Header.Get("Content-Type")

	return &Request{
		Header:      NewHeader(req.Header),
		Body:        NewBody(b2bytes, contentType),
		Method:      req.Method,
		Description: desc,
		Title:       title,
	}, nil
}

func (r *Request) Render() string {
	return render(requestTmpl, r)
}

func (r *Request) HasBody() bool {
	return r.Body != nil
}

func (r *Request) HasHeader() bool {
	return r.Header != nil && len(r.Header.DisplayHeader) > 0
}

func (r *Request) HasContentType() bool {
	return r.Header != nil && len(r.Header.ContentType) > 0
}
