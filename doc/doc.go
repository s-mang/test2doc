package doc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

const (
	FORMAT  = "1A"
	outFile = "apidoc.apib"
)

var (
	docTmpl *template.Template
	docFmt  = `FORMAT: {{with .Metadata}}{{.Format}}
HOST: {{.Host}}{{end}}

# {{.Title}}
{{.Description}}
{{range .ResourceGroups}}
{{.Render}}{{end}}`
)

func init() {
	docTmpl = template.Must(template.New("doc").Parse(docFmt))
}

type Doc struct {
	Title          string
	Description    string
	Metadata       Metadata
	ResourceGroups []*ResourceGroup
	file           *os.File

	// TODO:
	// DataStructures
}

type Metadata struct {
	Format string
	Host   string
}

func NewDoc(outDir string) (doc *Doc, err error) {
	var fi *os.File

	outPath := filepath.Join(outDir, outFile)
	fi, err = os.Create(outPath)
	if err != nil {
		return
	}

	doc = tmpDoc
	doc.file = fi

	return
}

// TODO: add Resource to appropriate ResourceGroup,
//  not just to ResourceGroups[0]
func (d *Doc) AddResource(resource *Resource) {
	group := d.ResourceGroups[0]
	group.Resources = append(group.Resources, *resource)
}

func (d *Doc) Write() error {
	return docTmpl.Execute(d.file, d)
}

func getPayload(req *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
