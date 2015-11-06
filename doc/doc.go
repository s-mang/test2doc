package doc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/adams-sarah/test2doc/doc/parse"
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

func NewDoc(pkgDir string) (doc *Doc, err error) {
	fiPath := filepath.Join(pkgDir, outFile)

	fi, err := os.Create(fiPath)
	if err != nil {
		return doc, err
	}

	pkgDoc, err := parse.NewPackageDoc(pkgDir)
	if err != nil {
		return doc, err

	} else if pkgDoc == nil {
		return doc, errors.New("Found 0 packages, expected 1.")
	}

	title, description, host := getDocInfo(pkgDoc.Doc)

	doc = &Doc{
		Title:       title,
		Description: description,
		Metadata: Metadata{
			Format: FORMAT,
			Host:   host,
		},
		ResourceGroups: tmpResourceGroups,

		file: fi,
	}

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

func getDocInfo(pkgDoc string) (title, description, host string) {
	pkgDocParts := strings.Split(pkgDoc, "\n")
	if len(pkgDocParts) > 0 {
		title = pkgDocParts[0]
	}

	if len(pkgDocParts) > 1 {
		description = pkgDocParts[1]
	}

	if len(pkgDocParts) > 2 {
		host = pkgDocParts[2]
	}

	return
}

func getPayload(req *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
