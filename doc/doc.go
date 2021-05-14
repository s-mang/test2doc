package doc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
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

type byResp []*Request

func (rs byResp) Len() int {
	return len(rs)
}
func (rs byResp) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
func (rs byResp) Less(i, j int) bool {
	if rs[i].Response.StatusCode != rs[j].Response.StatusCode {
		return rs[i].Response.StatusCode < rs[j].Response.StatusCode
	}
	var bodyIContentLen, bodyJContentLen int
	if rs[i].Response.Body != nil {
		bodyIContentLen = len(rs[i].Response.Body.Content)
	}
	if rs[j].Response.Body != nil {
		bodyIContentLen = len(rs[j].Response.Body.Content)
	}
	return bodyIContentLen > bodyJContentLen
}

// TODO: add Resource to appropriate ResourceGroup,
//  not just to ResourceGroups[0]
func (d *Doc) AddResource(resource *Resource) {
	// sort requests by response status code and body len
	for method, _ := range resource.Actions {
		sort.Sort(byResp(resource.Actions[method].Requests))
	}
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
