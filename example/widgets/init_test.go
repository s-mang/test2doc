package widgets

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/s-mang/prettytest"
	"github.com/s-mang/test2doc/test"
	"github.com/s-mang/test2doc/vars"
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
