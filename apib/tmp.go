package apib

// TODO: remove, TMP
type tmpSection struct {
	title string
	desc  string
}

var (
	tmpAPIBlueprint = tmpSection{
		title: "JSON Placeholder API",
		desc:  "Fake Online REST API for Testing and Prototyping",
	}
	tmpMetadataHost  = "http://jsonplaceholder.typicode.com"
	tmpResourceGroup = tmpSection{
		title: "A Lovely Resource Group",
		desc:  "All CRUD endpoints for A Lovely Resource Group",
	}
	tmpResource = tmpSection{
		title: "The Fanciest Resource",
		desc:  "This resource is just like all the others.",
	}
)
