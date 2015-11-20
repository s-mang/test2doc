package doc

import "text/template"

var (
	resourceTmpl *template.Template
	resourceFmt  = `## {{.URL.ParameterizedPath}}
{{.Description}}{{if .URL.Parameters}}
+ Parameters
{{range .URL.Parameters}}{{.Render}}
{{end}}{{end}}{{range .Actions}}
{{.Render}}{{end}}
`
)

func init() {
	resourceTmpl = template.Must(template.New("resource").Parse(resourceFmt))
}

type HTTPMethod string

type Resource struct {
	Title       string
	Description string
	//Model       *Model
	URL     *URL
	Actions map[HTTPMethod]*Action

	// TODO:
	// Attributes
}

// NewResource returns a new Resource object
func NewResource(u *URL) *Resource {
	resource := &Resource{}
	resource.Actions = map[HTTPMethod]*Action{}
	resource.URL = u
	return resource
}

func (r *Resource) AddAction(action *Action) {
	if r.Actions == nil {
		r.Actions = map[HTTPMethod]*Action{}
	}

	r.Actions[action.Method] = action
}

func (r *Resource) FindAction(httpMethod string) *Action {
	method := HTTPMethod(httpMethod)
	return r.Actions[method]
}

func (r *Resource) Render() string {
	return render(resourceTmpl, r)
}
