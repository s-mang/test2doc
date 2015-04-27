package blueprint

type ResourceGroup struct {
	Name        string
	Description string
	Resources   []*Resource
}

func NewResourceGroup(name, description string) *ResourceGroup {
	return &ResourceGroup{
		Name:        name,
		Description: description,
	}
}

func (group *ResourceGroup) AddResource(resource *Resource) {
	group.Resources = append(group.Resources, resource)
}
