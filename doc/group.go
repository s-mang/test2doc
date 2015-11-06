package doc

import "text/template"

var (
	resourceGroupTmpl *template.Template
	resourceGroupFmt  = `
# Group {{.Title}}
{{range .Resources}}
{{.Render}}{{end}}`
)

func init() {
	resourceGroupTmpl = template.Must(template.New("resourceGroup").Parse(resourceGroupFmt))
}

type ResourceGroup struct {
	Title       string
	Description string
	Resources   []Resource
}

func (rg *ResourceGroup) Render() string {
	return render(resourceGroupTmpl, rg)
}
