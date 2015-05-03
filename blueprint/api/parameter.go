package api

type ParameterType int

const (
	Number ParameterType = iota
	String
	Boolean
)

type Parameter struct {
	Name        string
	Description string
	Value       string
	Type        ParameterType
	IsRequired  bool

	// Todo:
	// DefaultValue
}
