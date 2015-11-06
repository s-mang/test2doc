package parse

import (
	"regexp"
	"strings"
)

const camelCase = "[A-Z]?[^A-Z]*"

// GetTitle extracts a title from the function name,
// where longFnName is of the form:
// github.com/adams-sarah/test2doc/example.HandleGetWidget
// and the out title would be:
// Handle Get Widget
func GetTitle(longFnName string) string {
	shortFnName := getShortFnName(longFnName)

	re := regexp.MustCompile(camelCase)

	words := re.FindAllString(shortFnName, -1)
	return strings.Join(words, " ")
}

func GetDescription(longFnName string) (desc string) {
	shortFnName := getShortFnName(longFnName)

	doc := funcsMap[shortFnName]
	if doc != nil {
		desc = strings.TrimPrefix(doc.Doc, shortFnName+" ")
	}

	return
}

// getShortFnName returns the name of the function, given
// longFnName of the form:
// github.com/adams-sarah/test2doc/example.HandleGetWidget
func getShortFnName(longFnName string) string {
	splitName := strings.Split(longFnName, ".")
	return splitName[len(splitName)-1]
}
