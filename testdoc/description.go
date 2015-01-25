package testdoc

import "text/template"

const APIBlueprintFormat = "1A"

type APIDescription struct {
	APIBlueprintFormat string
	HostURL            string
	Title              string
	Description        string
}

var descFmt = `FORMAT: {{.APIBlueprintFormat}}
HOST: {{.HostURL}}

# {{.Title}}
{{.Description}}

`

var descTmpl *template.Template

func init() {
	var err error
	descTmpl = template.New("APIDescription")

	descTmpl, err = descTmpl.Parse(descFmt)
	if err != nil {
		panic(err.Error)
	}
}

func (doc *APIDoc) WriteDescription() (err error) {
	if doc.desc == nil {
		doc.desc = &APIDescription{
			APIBlueprintFormat: APIBlueprintFormat,
			Title:              "API",
		}
	}

	return descTmpl.Execute(doc.file, doc.desc)
}
