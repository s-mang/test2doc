package api

import "net/url"

type URL struct {
	*url.URL
	Parameters []*Parameter
}

func (u *URL) PathWithParamNames() string {
	// TODO: replace params in path with {param-name}
	return u.Path
}
