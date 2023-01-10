package vars

import (
	"github.com/labstack/echo/v4"
	"github.com/s-mang/test2doc/doc/parse"
	"net/http"
	"sync"
)

// MakeEchoRouterExtractor is a URLVarExtractor for Echo server
func MakeEchoRouterExtractor(e *echo.Echo) parse.URLVarExtractor {
	router := e.Router()
	mu := sync.Mutex{}

	return func(req *http.Request) map[string]string {
		mu.Lock()
		defer mu.Unlock()
		if req == nil {
			panic("missing request")
		}

		c := e.AcquireContext()
		if c == nil {
			panic("cannot AcquireContext")
		}

		router.Find(req.Method, req.URL.Path, c)
		if c.Request() == nil {
			c.SetRequest(req)
		}
		params := c.QueryParams()
		c.Reset(nil, nil)
		e.ReleaseContext(c)

		out := make(map[string]string)
		for k, v := range params {
			out[k] = v[0]
		}

		return out
	}
}
