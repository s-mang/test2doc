package apib

import (
	"net/http"
	"net/http/httptest"
	"text/template"
)

var (
	responseTmpl *template.Template
	responseFmt  = `
+ Response {{.StatusCode}} {{with .ContentType}}({{.}}){{end}}
{{.Header.Render}}
{{.Body.Render}}`
)

func init() {
	responseTmpl = template.Must(template.New("response").Parse(responseFmt))
}

type Response struct {
	StatusCode  int
	Description string
	Header      http.Header
	Body        []byte

	// TODO:
	// Attributes
	// Schema
}

func NewResponse(description string, w *httptest.ResponseRecorder) (*Response, error) {

	body1, body2, err := cloneBody(w.Body)
	if err != nil {
		return nil, err
	}

	w.Body = body1

	return &Response{
		StatusCode:  w.Code,
		Description: description,
		Header:      w.Header(),
		Body:        body2.Bytes(),
	}, nil
}

func (r *Response) Render() string {
	return render(responseTmpl, r)
}

func (r *Response) ContentType() string {
	return r.Header.Get("Content-Type")
}

func (r *Response) BodyStr() string {
	return string(r.Body)
}
