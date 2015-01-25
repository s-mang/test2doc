package testdoc

import "text/template"

const (
	RequestKind  = "Request"
	ResponseKind = "Response"
)

type Title struct {
	Kind       string
	Properties PropertySet
}

var titleFmt = `+ {{.Kind}} {{.Properties.List}}

`

var titleTmpl *template.Template

func init() {
	var err error

	titleTmpl, err = template.New("Title").Parse(titleFmt)
	if err != nil {
		panic(err.Error)
	}
}

func (doc *APIDoc) WriteRequestTitle(description string) (err error) {
	title := Title{
		Kind: RequestKind,
		Properties: &RequestProperties{
			Description: description,
		},
	}
	return titleTmpl.Execute(doc.file, title)
}

func (doc *APIDoc) WriteResponseTitle(code int, contentType string) (err error) {
	title := Title{
		Kind: ResponseKind,
		Properties: &ResponseProperties{
			StatusCode:  code,
			ContentType: contentType,
		},
	}
	return titleTmpl.Execute(doc.file, title)
}
