package doc

import "text/template"

var (
	resourceTmpl *template.Template
	resourceFmt  = `## {{.Title}} [{{with .URL}}{{.ParameterizedPath}}{{end}}]
{{.Description}}
{{with .URL}}{{with .Parameters}}+ Parameters
{{range .}}
{{.Render}}{{end}}{{end}}{{end}}{{range .Actions}}
{{.Render}}{{end}}`
)

func init() {
	resourceTmpl = template.Must(template.New("resource").Parse(resourceFmt))
}

type Resource struct {
	Title       string
	Description string
	//Model       *Model
	URL     *URL
	Actions []*Action

	// TODO:
	// Attributes
}

func (r *Resource) Render() string {
	return render(resourceTmpl, r)
}
