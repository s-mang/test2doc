package doc

import (
	"fmt"
	"strings"
	"encoding/json"
	"net/http"
	"text/template"
)

var (
	requestTmpl *template.Template
	requestFmt  = `{{if or .HasBody .HasHeader}}
+ Request {{if .HasContentType}}({{.Header.ContentType}}){{end}}{{with .Header}}

{{.Render}}{{end}}{{if .Attributes}}
    + Attributes

{{range .Attributes}}{{.Render}}
{{end}}{{end}}{{with .Body}}
{{.Render}}{{end}}{{end}}`
)

func init() {
	requestTmpl = template.Must(template.New("request").Parse(requestFmt))
}

type Request struct {
	Header   *Header
	Body     *Body
	Method   string
	Response *Response

	Attributes []Attribute
	// Schema
}

func NewRequest(req *http.Request) (*Request, error) {
	body1, body2, err := cloneBody(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = nopCloser{body1}

	b2bytes := body2.Bytes()
	contentType := req.Header.Get("Content-Type")

	return &Request{
		Header: NewHeader(req.Header),
		Body:   NewBody(b2bytes, contentType),
		Method: req.Method,
		Attributes: getAttributesOf(contentType, b2bytes),
	}, nil
}

func getAttributesOf(contentType string, body []byte) []Attribute {
	var attrs []Attribute
	switch contentType {
	case "application/x-www-form-urlencoded":
		attrs = parseForm(body)
	case "application/json":
		attrs = parseJSON(body)
	}
	return attrs
}

func parseForm(body []byte) []Attribute {
	pairs := strings.Split(string(body[:]), "&")
	attrs := make([]Attribute, 0)
	for _, pair := range pairs {
		if len(pair) == 0 {
			continue
		}
		kv := strings.Split(pair, "=")
		if len(kv) <= 1 {
			continue
		}
		key, val := kv[0], kv[1]

		attr := attributeOf(key, val)
		attrs = append(attrs, attr)
	}
	return attrs
}

func attributeOf(key string, val interface{}) Attribute {
	s := fmt.Sprintf("%s", val)
	description, isRequired, defaultValue := getPropertyOf(key)
	return Attribute{
		Name: key,
		Description: description,
		Value: ParameterValue(s),
		Type:  paramType(s),
		IsRequired: isRequired,
		DefaultValue: defaultValue,
	}
}

func parseJSON(body []byte) (attrs []Attribute) {
	var obj interface{}
	err := json.Unmarshal(body, &obj)
	if err != nil {
		return
	}

	m := obj.(map[string]interface{})
	for key, val := range m {
		attr := attributeOf(key, val)
		attrs = append(attrs, attr)
	}
	return
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
