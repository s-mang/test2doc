package blueprint

import "github.com/adams-sarah/test2doc/blueprint/api"

const (
	FORMAT = "1A"
)

type APIBlueprint struct {
	Name           string
	Description    string
	Metadata       *Metadata
	ResourceGroups []*api.ResourceGroup

	// Todo:
	// DataStructures
}

type Metadata struct {
	Format string
	Host   string
}

func NewAPIBlueprint(name, desc, host string) *APIBlueprint {
	return &APIBlueprint{
		Name:        name,
		Description: desc,
		Metadata: &Metadata{
			Format: FORMAT,
			Host:   host,
		},
	}
}
