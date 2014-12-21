package middleware

import (
	"github.com/go-martini/martini"
	"net/http"
)

func Martini() interface{} {
	return func(c martini.Context, res http.ResponseWriter, req *http.Request) {
		c.Next()
	}
}