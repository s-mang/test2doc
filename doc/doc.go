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

var (
	docTmpl *template.Template
	docFmt  = `{{with .Group}}{{.Render}}{{end}}`
)

func init() {
	docTmpl = template.Must(template.New("doc").Parse(docFmt))
}

type Doc struct {
	Group ResourceGroup
	file  *os.File
}

type Metadata struct {
	Format string
	Host   string
}

func NewDoc(pkgDir string) (doc *Doc, err error) {
	pkgDoc, err := parse.NewPackageDoc(pkgDir)
	if err != nil {
		return doc, err

	} else if pkgDoc == nil {
		return doc, errors.New("Found 0 packages, expected 1.")
	}

	fiPath := filepath.Join(pkgDir, pkgDoc.Name+".apib")

	fi, err := os.Create(fiPath)
	if err != nil {
		return doc, err
	}

	doc = &Doc{
		Group: ResourceGroup{
			Title: strings.Title(pkgDoc.Name),
		},
		file: fi,
	}

	return
}

// TODO: add Resource to appropriate ResourceGroup,
//  not just to ResourceGroups[0]
func (d *Doc) AddResource(resource *Resource) {
	d.Group.Resources = append(d.Group.Resources, *resource)
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
