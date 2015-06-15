package doc

import "text/template"

var (
	resourceTmpl *template.Template
	resourceFmt  = `## {{.Title}} [{{.URL.ParameterizedPath}}]
{{.Description}}
{{if .URL.Parameters}}+ Parameters
{{range .URL.Parameters}}{{.Render}}{{end}}{{end}}
{{range .Actions}}{{.Render}}{{end}}
`
)

func init() {
	resourceTmpl = template.Must(template.New("resource").Parse(resourceFmt))
}

type httpMethod string

type Resource struct {
	Title       string
	Description string
	//Model       *Model
	URL     *URL
	Actions map[httpMethod]*Action

	// TODO:
	// Attributes
}

func NewResource(u *URL) *Resource {
	resource := &Resource{}
	resource.Actions = map[httpMethod]*Action{}
	resource.URL = u
	return resource
}

func (r *Resource) AddAction(action *Action) {
	if r.Actions == nil {
		r.Actions = map[httpMethod]*Action{}
	}

	r.Actions[action.Method] = action
}

func (r *Resource) Render() string {
	return render(resourceTmpl, r)
}
