package blueprint

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
