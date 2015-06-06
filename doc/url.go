package doc

import "net/url"

type URL struct {
	*url.URL
	Parameters []*Parameter
}

// TODO: replace params in path with {param-name}
func (u *URL) ParameterizedPath() string {
	return u.Path
}
