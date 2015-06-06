package parse

import "regexp"

type section int

const (
	blueprint section = iota
	resource
	action
)

var sectionRegexps = map[section]*regexp.Regexp{
	resource: regexp.MustCompile("^([A-Z][^ ]+) is a resource of the (.*)(?: API)?"),
	action:   regexp.MustCompile("^([A-Z][^ ]+) is API action for the (.*) resource"),
}

// func extract(doc string, re *regexp.Regexp)
