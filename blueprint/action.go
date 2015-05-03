package blueprint

import "github.com/adams-sarah/test2doc/blueprint/api"

type Action struct {
	Name        string
	Description string
	Parameters  []*api.Parameter
	Requests    []*api.Request
	Responses   []*api.Response

	// Todo:
	// Relation
	// Attributes
	// Schema
}
