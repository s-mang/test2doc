package blueprint

type ParameterType int

const (
	Number ParameterType = iota
	String
	Boolean
)

type Parameter struct {
	Name         string
	ExampleValue string
	Type         ParameterType
	IsRequired   bool
	Description  string

	// Todo:
	// - DefaultValue
}

func NewParameter(val string, t ParameterType, isRequired bool, description string) *Parameter {
	return &Parameter{
		ExampleValue: val,
		Type:         t,
		IsRequired:   isRequired,
		Description:  description,
	}
}
