package doc

import (
	"net/http"
	"net/url"
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
	u.ParameterizedPath = paramPath(req.URL.String())
	return u
}

// TODO: replace params in path with {param-name}
func paramPath(urlPath string) string {
	return urlPath
}
