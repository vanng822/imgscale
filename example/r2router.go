package main

import (
	"fmt"
	"github.com/vanng822/imgscale/imgscale"
	"github.com/vanng822/r2router"
	"net/http"
)

func main() {
	app := r2router.NewSeeforRouter()
	mhandler := imgscale.Configure("./config/formats.json")
	defer mhandler.Cleanup()
	app.Before(mhandler.Middleware())
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)
}
