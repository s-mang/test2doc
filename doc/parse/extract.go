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

// getShortFnName returns the name of the function
// without the package name so:
//   github.com/user/project/package.method
// becomes
//   method
// and
//   github.com/user/project/package.(*type).method
// becomes
//   type.method
func getShortFnName(longFnName string) string {

	// drop anything before the last '/'
	slashed := strings.Split(longFnName, "/")
	last := slashed[len(slashed)-1]

	// split the final part by period
	dotted := strings.Split(last, ".")

	// drop the first part which is the package name
	dotted = dotted[1:]

	// loop over and drop pointer references (*v) => v
	for i, p := range dotted {
		if len(p) > 3 {
			if p[0:2] == "(*" && p[len(p)-1] == ')' {
				p = p[2 : len(p)-1]
			}
		}
		dotted[i] = p
	}

	return strings.Join(dotted, ".")
}
