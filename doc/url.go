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
	uri, err := url.QueryUnescape(req.URL.Path)
	if err != nil {
		// fall back to unescaped uri
		uri = req.URL.Path
	}

	vars := (*parse.Extractor)(req)
	params := []Parameter{}

	for k, v := range vars {
		uri = strings.Replace(uri, "/"+v, "/{"+k+"}", 1)
		params = append(params, MakeParameter(k, v))
	}

	var queryKeys []string
	queryParams := req.URL.Query()

	for k, vs := range queryParams {
		queryKeys = append(queryKeys, k)

		// just take first value
		params = append(params, MakeParameter(k, vs[0]))
	}

	var queryKeysStr string
	if len(queryKeys) > 0 {
		queryKeysStr = "{?" + strings.Join(queryKeys, ",") + "}"
	}

	uri = uri + queryKeysStr

	return uri, params
}
