package doc

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	requestTmpl *template.Template
	requestFmt  = `{{if or .Header .Body}}
+ Request {{with .Header}}({{.ContentType}}){{end}}{{end}}
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

func (r *Request) ContentType() string {
	return r.Header.ContentType()
}

func (r *Request) BodyStr() string {
	fbody, err := formatBody(string(r.Body), r.ContentType())
	if err != nil {
		panic(err.Error())
	}

	return fbody
}

func RecordRequest(doc *Doc, req *http.Request) error {
	body, err := getPayload(req)
	if err != nil {
		return err
	}

	fmt.Println(body)

	// err = doc.WriteRequestTitle("")
	// if err != nil {
	// 	return err
	// }

	// err = doc.WriteHeaders(req.Header)
	// if err != nil {
	// 	return err
	// }

	// return doc.WriteBody(string(body))

	return nil
}
