package doc

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
	Header      Header
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
		Header:      Header(w.Header()),
		Body:        body2.Bytes(),
	}, nil
}

func (r *Response) Render() string {
	return render(responseTmpl, r)
}

func (r *Response) ContentType() string {
	return r.Header.ContentType()
}

func (r *Response) BodyStr() string {
	fbody, err := formatBody(string(r.Body), r.ContentType())
	if err != nil {
		panic(err.Error())
	}

	return fbody
}

func RecordResponse(doc *Doc, handler http.Handler, req *http.Request) (resp *httptest.ResponseRecorder, err error) {
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	// err = doc.WriteResponseTitle(resp.Code, resp.Header().Get("Content-Type"))
	// if err != nil {
	// 	return
	// }

	// err = doc.WriteHeaders(resp.Header())
	// if err != nil {
	// 	return
	// }

	// err = doc.WriteBody(string(resp.Body.String()))
	return
}
