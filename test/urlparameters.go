package test

import "regexp"

// Three general formats for URL params
// in common go routing packages:
// 1. "/hello/:name"  - httprouter, martini, goji, traffic
// 2. "/hello/{name}" - gorilla/mux
// 3. "/hello/#name"  - go-json-rest

const (
	curlyBracePattern = `/\{([^\}]+)\}`
	// colonPattern = ...
	// hashPattern = ...
)

var (
	CurlyBraceMatcher = regexp.MustCompile(curlyBracePattern)
)
