package main

import (
	"fmt"
	"github.com/vanng822/imgscale/imgscale"
	"net/http"
)

func main() {
	// http.Handler
	handler := imgscale.Configure("./config/formats.json")
	defer handler.Cleanup()
	http.Handle("/", handler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), nil)
}
