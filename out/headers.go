package out

import (
	"os"
	"text/template"
)

var headersTmpl = `+ Headers

	{{range $key, $vals := .Headers}}{{$key}}: {{$vals | commajoin}}
	{{end}}
`

func WriteHeaders(file *os.File, headers map[string][]string) (err error) {
	t := template.New("Headers template")

	t = t.Funcs(template.FuncMap{"commajoin": CommaJoinStrs})

	t, err = t.Parse(headersTmpl)
	if err != nil {
		return err
	}

	return t.Execute(file, map[string]interface{}{
		"Headers": headers,
	})
}
