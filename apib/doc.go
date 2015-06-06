package apib

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
	docFmt  = `FORMAT: {{.Metadata.Format}}
HOST: {{.Metadata.Host}}

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
	Metadata       *Metadata
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

	err = docTmpl.Execute(fi, doc)
	if err != nil {
		return
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
