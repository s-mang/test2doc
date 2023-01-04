package parse

import (
	"regexp"
	"strings"
)

const camelCase = "[A-Z]?[^A-Z]*"

// GetTitle extracts a title from the function name,
// where longFnName is of the form:
// github.com/adams-sarah/test2doc/example.GetWidget
// and the out title would be:
// Handle Get Widget
func GetTitle(longFnName string) string {
	shortFnName := getShortFnName(longFnName)

	re := regexp.MustCompile(camelCase)

	words := re.FindAllString(shortFnName, -1)

	for i := range words {
		words[i] = strings.Title(words[i])
	}

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

// IsFuncInPkg checks if this func belongs to the package
func IsFuncInPkg(longFnName string) bool {
	shortFnName := getShortFnName(longFnName)
	doc := funcsMap[shortFnName]
	return doc != nil
}

// getShortFnName returns the name of the function, given
// longFnName of the form:
// github.com/adams-sarah/test2doc/example.GetWidget
func getShortFnName(longFnName string) string {
	splitName := strings.Split(strings.Replace(longFnName, ").", ")", -1), ".")
	return splitName[len(splitName)-1]
}
