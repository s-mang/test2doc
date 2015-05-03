package blueprint

type ResourceGroup struct {
	Name        string
	Description string
	Resources   []*Resource
}
