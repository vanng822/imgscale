package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/go-martini/martini"
	"github.com/vanng822/imgscale/imgscale"
	"net/http"
)

func main() {
	n := negroni.New()
	handler := imgscale.Configure("./config/formats.json")
	defer handler.Cleanup()
	// Example how to run an arbitrary remote image provider
	handler.SetImageProvider(imgscale.NewImageProviderHTTP(""))
	n.UseHandler(handler)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8081), n)

	// Martini
	app := martini.Classic()
	mhandler := imgscale.Configure("./config/formats.json")
	defer mhandler.Cleanup()
	app.Use(mhandler.ServeHTTP)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)

	// http.Handler
	handler2 := imgscale.Configure("./config/formats.json")
	defer handler2.Cleanup()
	// Example how to run an host limited remote image provider, can not run arbitrary here
	handler2.SetImageProvider(imgscale.NewImageProviderHTTP("http://127.0.0.1:8080/img/original/"))
	http.Handle("/", handler2)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8082), nil)
}
