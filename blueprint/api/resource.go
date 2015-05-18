package api

type Resource struct {
	Name        string
	Description string
	Model       *Model
	URL         []*URL
	Actions     []*Action

	// Todo:
	// Attributes
}
