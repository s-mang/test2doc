package testdoc

import "text/template"

var bodyFmt = `	+ Body

			{{.JSONBody | indent}}

`

var bodyTmpl *template.Template

func init() {
	var err error
	bodyTmpl = template.New("Body").Funcs(template.FuncMap{"indent": IndentJSONBody})

	bodyTmpl, err = bodyTmpl.Parse(bodyFmt)
	if err != nil {
		panic(err.Error)
	}
}

func (doc *APIDoc) WriteBody(jsonBody string) (err error) {
	if len(jsonBody) == 0 {
		return
	}

	return bodyTmpl.Execute(doc.file, map[string]interface{}{
		"JSONBody": jsonBody,
	})
}
