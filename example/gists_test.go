package main

import (
	"net/http"
	"testing"

	"github.com/gophergala/test2doc/testdoc"
	"github.com/gophergala/test2doc/testrun"
)

type GistsGroup struct {
	testdoc.Group
}

func TestGistsGroup(t *testing.T) {
	gistsGroup := &GistsGroup{
		testdoc.Group{
			"Gists",
			"A Gist is a simple way to share snippets and pastes with others. All gists are Git repositories, so they are automatically versioned, forkable and usable from Git.",
		},
	}

	testrun.RunTests(t, docFile, gistsGroup)
}

func (g *GistsGroup) TestCompareGistsWithoutRequestBody(t *testing.T) {
	exp := &responseExpectation{
		StatusCode:   http.StatusBadRequest,
		BodyContains: []string{"No request body"},
	}

	err := helperTestGetRequest(server.URL+"/gists/1/compare", exp)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func (g *GistsGroup) TestCompareGists(t *testing.T) {
	reqBody := `{"Id": 2, "Title": "Some other Gist"}`
	exp := &responseExpectation{
		StatusCode:   http.StatusOK,
		BodyContains: []string{"Diffs"},
	}

	helperTestPostRequest(server.URL+"/gists/1/compare", reqBody, exp)
}

func (g *GistsGroup) TestGetGist(t *testing.T) {
	exp := &responseExpectation{
		StatusCode:      http.StatusOK,
		BodyContains:    []string{"MyGist", "1"},
		BodyNotContains: []string{"Some other Gist", "2"},
	}

	err := helperTestGetRequest(server.URL+"/gists/1", exp)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func (g *GistsGroup) TestGetGists(t *testing.T) {
	exp := &responseExpectation{
		StatusCode:   http.StatusOK,
		BodyContains: []string{"MyGist", "1", "Some other Gist", "2"},
	}

	err := helperTestGetRequest(server.URL+"/gists", exp)
	if err != nil {
		t.Fatal(err.Error())
	}
}
