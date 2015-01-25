package testrun

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/gophergala/test2doc/testdoc"
)

var testFnRegxp *regexp.Regexp

func init() {
	testFnRegxp = regexp.MustCompile("^Test")
}

func RunTests(t *testing.T, doc *testdoc.APIDoc, runGroup testdoc.TestGroup) {
	err := doc.WriteGroup(runGroup)
	if err != nil {
		panic(err.Error())
	}

	groupType := reflect.TypeOf(runGroup)
	for i := 0; i < groupType.NumMethod(); i++ {
		if testFnRegxp.MatchString(groupType.Method(i).Name) {
			reflect.ValueOf(runGroup).Method(i).Call(
				[]reflect.Value{
					reflect.ValueOf(t),
				},
			)
		}
	}
}
