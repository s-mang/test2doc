package main

import (
	"net/http"
	"testing"

	"github.com/gophergala/test2doc/testdoc"
	"github.com/gophergala/test2doc/testrun"
)

type ReposGroup struct {
	testdoc.Group
}

func TestReposGroup(t *testing.T) {
	reposGroup := &ReposGroup{
		testdoc.Group{
			"Repos",
			"A Repo is an on-disk data structure which stores metadata for a set of files and/or directory structure.",
		},
	}

	testrun.RunTests(t, docFile, reposGroup)
}

func (g *ReposGroup) TestGetRepo(t *testing.T) {
	exp := &responseExpectation{
		StatusCode:      http.StatusOK,
		BodyContains:    []string{"my_repo", "1", "adams-sarah"},
		BodyNotContains: []string{"my_other_repo", "someone_elses_repo"},
	}

	err := helperTestGetRequest(server.URL+"/repos/1", exp)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func (g *ReposGroup) TestGetRepos(t *testing.T) {
	exp := &responseExpectation{
		StatusCode:   http.StatusOK,
		BodyContains: []string{"my_repo", "1", "adams-sarah", "my_other_repo", "someone_elses_repo", "2", "3"},
	}

	err := helperTestGetRequest(server.URL+"/repos", exp)
	if err != nil {
		t.Fatal(err.Error())
	}
}
