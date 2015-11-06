package parse

import "net/http"

type URLVarExtractor func(req *http.Request) (vars map[string]string)

var Extractor *URLVarExtractor

func SetURLVarExtractor(fn *URLVarExtractor) {
	Extractor = fn
}
