package doc

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"reflect"
	"strings"
	"text/template"

	"github.com/everytv/test2doc/doc/parse"
)

var (
	responseTmpl *template.Template
	responseFmt  = `+ Response {{.StatusCode}} {{if .HasContentType}}({{.Header.ContentType}}){{end}}{{with .Header}}

{{.Render}}{{end}}
    + Attributes ({{.DataStructure.Name}})

{{with .DataStructure}}
{{.Render}}{{end}}
{{with .Body}}
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
	// Schema
	DataStructure *DataStructure
}

func NewResponse(resp *httptest.ResponseRecorder) *Response {
	content := resp.Body.Bytes()
	contentType := resp.Header().Get("Content-Type")
	dataStructure := getAttributesOfResponse(contentType, content)

	return &Response{
		StatusCode:    resp.Code,
		Header:        NewHeader(resp.Header()),
		Body:          NewBody(content, contentType),
		DataStructure: dataStructure,
	}
}

func (r *Response) Render() string {
	return render(responseTmpl, r)
}

func (r *Response) HasContentType() bool {
	return r.Header != nil && len(r.Header.ContentType) > 0
}

func getAttributesOfResponse(contentType string, body []byte) *DataStructure {
	if len(contentType) == 0 {
		return nil
	}

	var response map[string]interface{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}

	data := response["data"]
	var attrs []Attribute
	for k, v := range data.(map[string]interface{}) {
		value := reflect.Indirect(reflect.ValueOf(v))
		t := value.Type()

		var isObject bool
		s := fmt.Sprintf("%s", v)
		if t.Kind() == reflect.Struct || t.Kind() == reflect.Slice || t.Kind() == reflect.Map || t.Kind() == reflect.Ptr {
			isObject = true
			s = ""
			// TODO: 型の名前を取得しなければいけない（publishersでいうと、Advertisersという文字を取りたい）
		}

		description, isRequired, defaultValue := getPropertyOfResponse(k)
		attr := Attribute{
			Name:         k,
			Description:  description,
			Value:        ParameterValue(s),
			Type:         paramType(s),
			IsRequired:   isRequired,
			DefaultValue: defaultValue,
			IsObject:     isObject,
		}
		attrs = append(attrs, attr)
	}

	responsePackageStr := fmt.Sprintf("%s", *parse.ResponseType)
	responseStructName := strings.Split(responsePackageStr, ".")[1]
	return &DataStructure{
		Name:       responseStructName + "Response",
		ObjectName: responseStructName,
		Attributes: attrs,
	}
}
