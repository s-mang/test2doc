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

func (group *ResourceGroup) SetParameters(resource *Resource) {
	group.Resources = append(group.Resources, resource)
}
