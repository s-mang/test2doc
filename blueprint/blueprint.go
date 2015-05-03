package blueprint

const (
	FORMAT = "1A"
)

type APIBlueprint struct {
	Name           string
	Description    string
	Metadata       *Metadata
	ResourceGroups []*ResourceGroup

	// Todo:
	// DataStructures
}

type Metadata struct {
	Format string
	Host   string
}
