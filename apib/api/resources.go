package api

type ResourceGroup struct {
	Name        string
	Description string
	Resources   []*Resource
}

type Resource struct {
	Name        string
	Description string
	//Model       *Model
	URL     *URL
	Actions *[]Action

	// TODO:
	// Attributes
}
