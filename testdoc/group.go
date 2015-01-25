package testdoc

import "text/template"

// A Group is a group of API endpoints for one resource
// 'Group' is an API Blueprint keyword.

type Group struct {
	Title       string // eg. Gists
	Description string
	//BaseURIs    BaseURISet
}

type BaseURISet struct {
	CollectionRoute string // eg. /gists
	SingularRoute   string // eg. /gists/{id}
}

type TestGroup interface {
	GetTitle() string
	GetDescription() string
}

func (g *Group) GetTitle() string {
	return g.Title
}

func (g *Group) GetDescription() string {
	return g.Description
}

var groupFmt = `# Group {{.GetTitle}}
{{.GetDescription}}

`

var groupTmpl *template.Template

func init() {
	var err error
	groupTmpl = template.New("Group")

	groupTmpl, err = groupTmpl.Parse(groupFmt)
	if err != nil {
		panic(err.Error)
	}
}

func (doc *APIDoc) WriteGroup(group TestGroup) (err error) {
	return groupTmpl.Execute(doc.file, group)
}

/*

USAGE THOUGHTS:

type GistGroup struct {
	testdoc.Group
}

func TestGistGroup(t *testing.T) {
	gistsGroup := &GistGroup{
		"Gists",
		"A Gist is a simple way to share snippets and pastes with others. All gists are Git repositories, so they are automatically versioned, forkable and usable from Git.",
	}

	testdoc.RunTests(t, gistsGroup)
}

func (g *GistsGroup) TestRetrieveASingleGist(t *testing.T) {
	// ...testing
}

func (g *GistsGroup) TestRetrieveAllGists(t *testing.T) {
	// ...testing
}

*/
