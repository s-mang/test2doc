package doc

var (
	tmpDoc = &Doc{
		Title:       "JSON XYZ API",
		Description: "Fake Online REST API for Testing and Prototyping",
		Metadata: Metadata{
			Format: FORMAT,
			Host:   "http://jsonplaceholder.typicode.com",
		},
		ResourceGroups: []*ResourceGroup{
			&ResourceGroup{
				Title:       "A Lovely Resource Group",
				Description: "All CRUD endpoints for A Lovely Resource Group",
			},
		},
	}
)
