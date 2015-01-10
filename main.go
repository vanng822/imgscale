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
	nh := imgscale.Configure("./config/formats.json")
	n.UseHandler(nh)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8081), n)

	// Martini
	app := martini.Classic()
	app.Use(imgscale.Configure("./config/formats.json").HandleFunc)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)

	// http.Handler
	http.Handle("/", imgscale.Configure("./config/formats.json"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8082), nil)
}
