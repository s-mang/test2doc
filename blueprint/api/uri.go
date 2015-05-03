package blueprint

type URI struct {
	Raw        string
	Tmpl       string
	Parameters []*Parameter
}

// type ParamMatcherFn func(str string) []*Parameter

// var ParamMatcher = defaultParamMatcherFn

// // paramRegexp is for gorilla/mux-like uri parameters
// // eg. "/articles/{category}/{id:[0-9]+}"
// var defaultParamRegexp = regexp.MustCompile(`\{[^\:\}]+(\:[^\}]+)?\}`)

// func defaultParamMatcherFn(str string) []*Parameter {

// }
