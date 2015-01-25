package testdoc

import "text/template"

var headersBlacklist = map[string]bool{
	"Accept":          true,
	"Accept-Encoding": true,
	"Accept-Language": true,
	"Connection":      true,
	"Content-Length":  true,
	"Content-Type":    true,
	"User-Agent":      true,
}

var headersFmt = `	+ Headers

			{{range $key, $vals := .Headers}}{{$key}}: {{$vals | commajoin}}
			{{end}}

`

var headersTmpl *template.Template

func init() {
	var err error
	headersTmpl = template.New("Headers").Funcs(template.FuncMap{"commajoin": CommaJoinStrs})

	headersTmpl, err = headersTmpl.Parse(headersFmt)
	if err != nil {
		panic(err.Error)
	}
}

func (doc *APIDoc) WriteHeaders(headers map[string][]string) (err error) {
	for k, _ := range headers {
		if headersBlacklist[k] {
			delete(headers, k)
		}
	}

	if len(headers) == 0 {
		return
	}

	return headersTmpl.Execute(doc.file, map[string]interface{}{
		"Headers": headers,
	})
}
