package test

import "github.com/happyreturns/test2doc/doc/parse"

func RegisterURLVarExtractor(fn parse.URLVarExtractor) {
	parse.SetURLVarExtractor(&fn)
}
