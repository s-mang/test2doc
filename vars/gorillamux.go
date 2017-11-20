package vars

import (
	"net/http"

	"github.com/adams-sarah/test2doc/doc/parse"
	"github.com/gorilla/mux"
)

func MakeGorillaMuxExtractor(router *mux.Router) parse.URLVarExtractor {
	return func(req *http.Request) map[string]string {
		// We must perform the match ourselves, as
		// context is cleared after the request has been handled,
		// and the vars are not set for mux.Vars until ServeHTTP is
		// called.
		var match mux.RouteMatch
		if router.Match(req, &match) {
			return match.Vars
		}

		return nil
	}
}
