package vars

import (
	"net/http"
	"strings"

	"github.com/happyreturns/test2doc/doc/parse"
	"github.com/julienschmidt/httprouter"
)

func MakeHTTPRouterExtractor(router *httprouter.Router) parse.URLVarExtractor {
	return func(req *http.Request) map[string]string {
		// httprouter Lookup func needs a trailing slash on path
		path := req.URL.Path
		if !strings.HasSuffix(path, "/") {
			path += "/"
		}

		_, params, ok := router.Lookup(req.Method, path)
		if !ok {
			return nil
		}

		paramsMap := make(map[string]string, len(params))
		for _, p := range params {
			paramsMap[p.Key] = p.Value
		}

		return paramsMap
	}
}
