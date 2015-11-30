package doc

import (
	"net/http"
	"text/template"
)

var (
	headerTmpl *template.Template
	headerFmt  = `{{with .DisplayHeader}}    + Headers
            {{range $name, $vals := .}}
            {{$name}}: {{$vals | commaJoin}}{{end}}{{end}}
`
)

func init() {
	funcMap := template.FuncMap{
		"commaJoin": commaJoin,
	}

	headerTmpl = template.Must(template.New("headers").Funcs(funcMap).Parse(headerFmt))
}

type Header struct {
	DisplayHeader http.Header
	ContentType   string
}

func NewHeader(h http.Header) *Header {
	hCopy := make(http.Header, len(h))
	CopyHeader(hCopy, h)

	// remove header fields we don't want in the doc
	delete(hCopy, "Accept-Encoding")
	delete(hCopy, "User-Agent")
	delete(hCopy, "Content-Length")

	if len(hCopy) == 0 {
		return nil
	}

	contentType := hCopy.Get("Content-Type")
	delete(hCopy, "Content-Type")

	return &Header{
		DisplayHeader: hCopy,
		ContentType:   contentType,
	}
}

func (h Header) Render() string {
	return render(headerTmpl, h)
}
