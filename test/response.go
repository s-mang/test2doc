package test

import (
	"reflect"

	"github.com/everytv/test2doc/doc/parse"
)

func RegisterResponseType(responseType *reflect.Type) {
	parse.SetResponseType(responseType)
}
