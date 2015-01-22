package main


import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/vanng822/imgscale/imgscale"
	"net/http"
)

func main() {
	n := negroni.New()
	handler := imgscale.Configure("./config/formats.json")
	defer handler.Cleanup()
	// Example how to run an arbitrary remote image provider
	n.UseHandler(handler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), n)
}
