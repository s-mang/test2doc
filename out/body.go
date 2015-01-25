package out

import (
	"os"
	"text/template"
)

var bodyFmt = `	+ Body

			{{.JSONBody | indent}}

`

var bodyTmpl *template.Template

func init() {
	var err error
	bodyTmpl = template.New("Headers").Funcs(template.FuncMap{"indent": IndentJSONBody})

	bodyTmpl, err = bodyTmpl.Parse(bodyFmt)
	if err != nil {
		panic(err.Error)
	}
}

func WriteBody(file *os.File, jsonBody string) (err error) {
	return bodyTmpl.Execute(file, map[string]interface{}{
		"JSONBody": jsonBody,
	})
}
