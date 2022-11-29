package widgets

import (
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/s-mang/test2doc/test"
	"github.com/s-mang/test2doc/vars"
)

var router *mux.Router
var server *test.Server

func TestMain(m *testing.M) {
	var err error

	router = mux.NewRouter()
	AddRoutes(router)

	test.RegisterURLVarExtractor(vars.MakeGorillaMuxExtractor(router))

	server, err = test.NewServer(router)
	if err != nil {
		panic(err.Error())
	}
	code := m.Run()
	server.Finish()
	os.Exit(code)
}
