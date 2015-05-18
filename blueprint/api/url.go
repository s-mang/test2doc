package api

import "net/url"

type URL struct {
	*url.URL
	Parameters []*Parameter
}
