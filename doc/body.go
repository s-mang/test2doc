package doc

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"
)

var (
	bodyTmpl *template.Template
	bodyFmt  = `    + Body

            {{.FormattedStr}}        
`
)

var multipartBoundaryREStr = "multipart/form-data; boundary=([-]*[a-zA-Z0-9]+)"
var multipartFileREStr = "(Content-Disposition: .*filename=.*\n?(?:Content-Type: .*)?)"

var multipartBoundaryRE, multipartFileRE *regexp.Regexp

func init() {
	bodyTmpl = template.Must(template.New("body").Parse(bodyFmt))

	multipartBoundaryRE = regexp.MustCompile(multipartBoundaryREStr)
	multipartFileRE = regexp.MustCompile(multipartFileREStr)
}

type Body struct {
	Content     []byte
	ContentType string
}

func NewBody(content []byte, contentType string) (b *Body) {
	if len(content) > 0 {
		b = &Body{
			Content:     content,
			ContentType: contentType,
		}
	}

	return b
}

func (b *Body) Render() string {
	return render(bodyTmpl, b)
}

func (b *Body) FormattedStr() string {
	if strings.HasPrefix(b.ContentType, "application/json") {
		return b.FormattedJSON()
	}
	if strings.HasPrefix(b.ContentType, "multipart/form-data") {
		return b.SanitizedMultipartForm()
	}
	return string(b.Content)
}

func (b *Body) FormattedJSON() string {
	fbody, err := indentJSONBody(string(b.Content))
	if err != nil {
		panic(err.Error())
	}

	return fbody
}

func (b *Body) SanitizedMultipartForm() string {
	matches := multipartBoundaryRE.FindStringSubmatch(b.ContentType)
	if len(matches) < 2 {
		// Fail, just return full body
		return string(b.Content)
	}
	boundary := matches[1]
	parts := bytes.Split(b.Content, []byte(boundary))

	for i, p := range parts {
		fileMatches := multipartFileRE.FindSubmatch(p)
		if len(fileMatches) > 0 {
			parts[i] = append([]byte("\n"), fileMatches[0]...)
			parts[i] = append(parts[i], []byte("\n\n<FILE DATA>\n\n")...)
		}
	}

	return string(bytes.Join(parts, []byte(boundary)))
}
