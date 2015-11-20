package doc

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/adams-sarah/test2doc/doc/parse"
)

type URL struct {
	rawURL            *url.URL
	ParameterizedPath string
	Parameters        []Parameter
}

func NewURL(req *http.Request) *URL {
	u := &URL{
		rawURL: req.URL,
	}
	u.ParameterizedPath, u.Parameters = paramPath(req)
	return u
}

func paramPath(req *http.Request) (string, []Parameter) {
	uri, err := url.QueryUnescape(req.URL.String())
	if err != nil {
		// fall back to unescaped uri
		uri = req.URL.String()
	}

	vars := (*parse.Extractor)(req)
	params := []Parameter{}

	for k, v := range vars {
		uri = strings.Replace(uri, "/"+v, "/{"+k+"}", 1)
		params = append(params, MakeParameter(k, v))
	}

	return uri, params
}
