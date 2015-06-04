package apib

// TODO: remove, TMP
type tmpSection struct {
	title string
	desc  string
}

var (
	tmpDoc = &Doc{
		Title:       "JSON Placeholder API",
		Description: "Fake Online REST API for Testing and Prototyping",
		Metadata: &Metadata{
			Format: FORMAT,
			Host:   "http://jsonplaceholder.typicode.com",
		},
		ResourceGroups: []*ResourceGroup{
			&ResourceGroup{
				Title:       "A Lovely Resource Group",
				Description: "All CRUD endpoints for A Lovely Resource Group",
			},
			&ResourceGroup{
				Title:       "An Awesome Resource Group",
				Description: "All non-CRUD endpoints",
			},
		},
	}

	tmpResource = Resource{
		Title:       "The Fanciest Resource",
		Description: "This resource is just like all the others.",
	}
)
