package blueprint

import "github.com/adams-sarah/test2doc/blueprint/api"

type Resource struct {
	Name        string
	Description string
	Parameters  []*api.Parameter
	Actions     []*Action

	// Todo:
	// Attributes
	// Model
}
