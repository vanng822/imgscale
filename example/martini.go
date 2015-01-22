package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/vanng822/imgscale/imgscale"
	"net/http"
)

func main() {
	// Martini
	app := martini.Classic()
	mhandler := imgscale.Configure("./config/formats.json")
	defer mhandler.Cleanup()
	app.Use(mhandler.ServeHTTP)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)
}
