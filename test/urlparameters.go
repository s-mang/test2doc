package test

import "github.com/adams-sarah/test2doc/doc/parse"
import "reflect"

func RegisterURLVarExtractor(fn parse.URLVarExtractor) {
	parse.SetURLVarExtractor(&fn)
}

func RegisterParamsType(paramsType *reflect.Type) {
	parse.SetParamsType(paramsType)
}
