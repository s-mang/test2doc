package apib

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func cloneBody(r io.Reader) (*bytes.Buffer, *bytes.Buffer, error) {
	var clone1, clone2 bytes.Buffer

	rBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return &clone1, &clone2, err
	}

	mw := io.MultiWriter(&clone1, &clone2)
	_, err = mw.Write(rBytes)

	return &clone1, &clone2, err
}

func render(tmpl *template.Template, i interface{}) string {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, i)
	if err != nil {
		panic(err.Error())
	}

	return buf.String()
}

func commaJoin(args ...interface{}) string {
	var strList []string

	for _, arg := range args {
		strs, ok := arg.([]string)
		if ok {
			for _, str := range strs {
				strList = append(strList, str)
			}
		} else {
			log.Println("Error: CommaJoinStrs called with non []string argument.")
		}

	}

	return strings.Join(strList, ", ")
}

func indentJSONBody(bodyStr string) (outStr string, err error) {
	var outJSON bytes.Buffer
	err = json.Indent(&outJSON, []byte(bodyStr), "\t\t\t", "\t")
	if err != nil {
		return
	}

	return string(outJSON.Bytes()), nil
}
