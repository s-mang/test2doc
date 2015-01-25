package out

import (
	"os"
	"text/template"
)

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

func WriteHeaders(file *os.File, headers map[string][]string) (err error) {
	return headersTmpl.Execute(file, map[string]interface{}{
		"Headers": headers,
	})
}
