package test

import "github.com/adams-sarah/test2doc/doc/parse"

func RegisterURLVarExtractor(fn parse.URLVarExtractor) {
	parse.SetURLVarExtractor(&fn)
}
