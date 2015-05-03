package testrun

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/adams-sarah/test2doc/testdoc"
)

var testFnRegxp *regexp.Regexp

func init() {
	testFnRegxp = regexp.MustCompile("^Test")
}

func RunTests(t *testing.T, doc *testdoc.APIBlueprint, runResourceGroup testdoc.TestResourceGroup) {
	err := doc.WriteResourceGroup(runResourceGroup)
	if err != nil {
		panic(err.Error())
	}

	groupType := reflect.TypeOf(runResourceGroup)
	for i := 0; i < groupType.NumMethod(); i++ {
		if testFnRegxp.MatchString(groupType.Method(i).Name) {
			reflect.ValueOf(runResourceGroup).Method(i).Call(
				[]reflect.Value{
					reflect.ValueOf(t),
				},
			)
		}
	}
}
