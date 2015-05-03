package blueprint

type Resource struct {
	Name        string
	Description string
	Parameters  []*api.Parameter
	Actions     []*Action

	// Todo:
	// Attributes
	// Model
}
