package parse

import "net/http"
import "reflect"

type URLVarExtractor func(req *http.Request) (vars map[string]string)

var Extractor *URLVarExtractor

func SetURLVarExtractor(fn *URLVarExtractor) {
	Extractor = fn
}

var ParamsType *reflect.Type

func SetParamsType(paramsType *reflect.Type) {
	ParamsType = paramsType
}
