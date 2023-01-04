package vars

import (
	"github.com/labstack/echo/v4"
	"github.com/s-mang/test2doc/doc/parse"
	"net/http"
)

// MakeEchoRouterExtractor is a URLVarExtractor for Echo server
func MakeEchoRouterExtractor(e *echo.Echo) parse.URLVarExtractor {
	router := e.Router()
	return func(req *http.Request) map[string]string {
		out := make(map[string]string)
		c := e.AcquireContext()
		router.Find(req.Method, req.URL.Path, c)
		params := c.QueryParams()
		c.Reset(nil, nil)
		e.ReleaseContext(c)

		for k, v := range params {
			out[k] = v[0]
		}

		return out
	}
}
