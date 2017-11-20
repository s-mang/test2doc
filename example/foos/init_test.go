package foos

import (
	"testing"

	"github.com/adams-sarah/prettytest"
	"github.com/adams-sarah/test2doc/test"
	"github.com/adams-sarah/test2doc/vars"
	"github.com/gorilla/mux"
)

var router *mux.Router
var server *test.Server

type mainSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	var err error

	router = mux.NewRouter()
	AddRoutes(router)
	test.RegisterURLVarExtractor(vars.MakeGorillaMuxExtractor(router))

	server, err = test.NewServer(router)
	if err != nil {
		panic(err.Error())
	}
	defer server.Finish()

	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(mainSuite),
	)
}
