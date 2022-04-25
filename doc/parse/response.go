package parse

import "reflect"

var ResponseType *reflect.Type

func SetResponseType(responseType *reflect.Type) {
	ResponseType = responseType
}
