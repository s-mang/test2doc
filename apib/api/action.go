package api

type Action struct {
	Name        string
	Description string
	HTTPMethod  string
	Request     *Request // status OK
	Response    *Response

	// TODO: document non-OK requests ??
}
