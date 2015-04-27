package blueprint

const (
	FORMAT = "1A"
)

type APIBlueprint struct {
	Metadata       *Metadata
	Name           string
	Description    string
	ResourceGroups []*ResourceGroup

	// Todo:
	// - DataStructures
}

func NewAPIBlueprint(name, description string, host string) *APIBlueprint {
	return &APIBlueprint{
		Metadata:    NewMetadata(host),
		Name:        name,
		Description: description,
	}

}

type Metadata struct {
	Format string
	Host   string
}

func NewMetadata(host string) *Metadata {
	return &Metadata{
		Format: FORMAT,
		Host:   host,
	}
}
