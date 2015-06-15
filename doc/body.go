package doc

import "text/template"

var (
	bodyTmpl *template.Template
	bodyFmt  = `	+ Body

            {{.FormattedJSON}}        
`
)

func init() {
	bodyTmpl = template.Must(template.New("body").Parse(bodyFmt))
}

type Body []byte

func (b *Body) Render() string {
	return render(bodyTmpl, b)
}

// TODO: support other content-types
func (b *Body) FormattedJSON() string {
	fbody, err := indentJSONBody(string(*b))
	if err != nil {
		panic(err.Error())
	}

	return fbody
}
