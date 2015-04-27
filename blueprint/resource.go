package blueprint

type Resource struct {
	Name        string
	Description string
	Parameters  []*Parameter
	Actions     []*Action

	// Todo:
	// - Attributes
	// - Model
}

func NewResource(name, description string) *Resource {
	return &Resource{
		Name:        name,
		Description: description,
	}
}

func (group *ResourceGroup) SetParameters(resource *Resource) {
	group.Resources = append(group.Resources, resource)
}
