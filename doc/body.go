package doc

import (
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

var multipartBoundaryREStr = "^([-]+[a-zA-Z0-9]+)\nContent-Disposition.*"
var multipartFileREStr = "(Content-Disposition: .*filename=.*\n?(?:Content-Type: .*))\n\n"

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
	bodyStr := string(b.Content)
	matches := multipartBoundaryRE.FindStringSubmatch(bodyStr)
	if len(matches) < 2 {
		// Fail, just return full body
		return string(b.Content)
	}
	boundary := matches[1]
	parts := strings.Split(bodyStr, boundary+"\n")

	for i, p := range parts {
		fileMatches := multipartFileRE.FindStringSubmatch(p)
		if len(fileMatches) > 0 {
			parts[i] = fileMatches[0] + "<FILE DATA>\n\n"
		}
	}

	return strings.Join(parts, boundary+"\n") + boundary + "--"
}
