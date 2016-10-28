package doc

import (
	"strings"
	"text/template"
)

var (
	bodyTmpl *template.Template
	bodyFmt  = `    + Body

            {{.FormattedStr}}        
`
)

func init() {
	bodyTmpl = template.Must(template.New("body").Parse(bodyFmt))
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
	return string(b.Content)
}

func (b *Body) FormattedJSON() string {
	fbody, err := indentJSONBody(string(b.Content))
	if err != nil {
		panic(err.Error())
	}

	return fbody
}
