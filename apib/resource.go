package apib

import "text/template"

var (
	resourceTmpl *template.Template
	resourceFmt  = `## {{.Title}} [{{.URL.ParemeterizedPath}}]
{{.Description}}

{{with .URL.Parameters}}
+ Parameters
{{range .}}
{{.Render}}
{{end}}{{end}}

{{range .Actions}}
{{.Render}}
{{end}}
`
)

func init() {
	resourceTmpl = template.Must(template.New("resource").Parse(resourceFmt))
}

type Resource struct {
	Title       string
	Description string
	//Model       *Model
	URL     *URL
	Actions *[]Action

	// TODO:
	// Attributes
}

func (r *Resource) Render() string {
	return render(resourceTmpl, r)
}
