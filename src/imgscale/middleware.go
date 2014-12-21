package imgscale

import (
	"fmt"
	"net/http"
)

func Middleware(config *Config) func(res http.ResponseWriter, req *http.Request) {
	handler := Configure(config)
	fmt.Println(handler.Path)
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" && req.Method != "HEAD" {
			return
		}
		matched, info := handler.Match(req.URL.RequestURI())
		if !matched {
			return
		}
		handler.Serve(res, req, info)
	}
}
