package apib

import (
	"path/filepath"
	"strings"
	"text/template"
)

var docTemplate *template.Template

func init() {
	filenames, err := filepath.Glob("templates/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	for _, fname := range filenames {
		base := filepath.Base(fname)
		tmplName := strings.Split(base, ".")[0]
		tmpl := template.New(tmplName)
		tmpl, err = template.ParseFiles(fname)
		if err != nil {
			panic(err.Error())
		}

		if tmplName == "index" {
			docTemplate = tmpl
		}
	}
}
