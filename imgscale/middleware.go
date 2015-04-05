package imgscale

import (
	"github.com/vanng822/imgscale/imagick"
	"net/http"
)

/*
	Configure returns Handler which implement http.Handler
	Filename is the configuration file in json, content looks something like this

		{
			"Path": "./data",
			"Prefix": "img",
			"Formats": [
				{"Prefix": "100x100", "Height": 100, "Ratio": 1.0, "Thumbnail": true},
				{"Prefix": "66x100", "Height": 100, "Ratio": 0.67, "Thumbnail": true},
				{"Prefix": "100x75", "Height": 75, "Ratio": 1.335, "Thumbnail": true},
				{"Prefix": "100x0", "Height": 100, "Ratio": 0.0, "Thumbnail": true, "Watermark": true},
				{"Prefix": "originalx1", "Height": 0, "Ratio": 1.0, "Thumbnail": false, "Watermark": true},
				{"Prefix": "original", "Height": 0, "Ratio": 0.0, "Thumbnail": false, "Watermark": true}
			],
			"Separator": "/",
			"Exts": ["jpg", "jpeg", "png"],
			"Comment": "Copyright",
			"AutoRotate": true,
			"Watermark": {"Filename": "./data/eyes.gif"}
		}


	The returned handler could use as middleware handler

	Negroni middleware:

		app := negroni.New()
		handler := imgscale.Configure("./config/formats.json")
		defer handler.Cleanup()
		app.UseHandler(handler)
		http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)

	Martini middleware:

		app := martini.Classic()
		handler := imgscale.Configure("./config/formats.json")
		defer handler.Cleanup()
		app.Use(handler.ServeHTTP)
		http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)

	http.Handle:

		handler := imgscale.Configure("./config/formats.json")
		defer handler.Cleanup()
		http.Handle("/", handler)
		http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8080), nil)

*/
func Configure(filename string) Handler {
	imagick.Initialize()
	defer func() {
		if r := recover(); r != nil {
			imagick.Terminate()
			panic(r)
		}
	}()
	config := LoadConfig(filename)
	return configure(config)
}

func Middleware(filename string) func(next http.Handler) http.Handler {
	imagick.Initialize()
	defer func() {
		if r := recover(); r != nil {
			imagick.Terminate()
			panic(r)
		}
	}()
	config := LoadConfig(filename)
	handler := configure(config)
	return func(next http.Handler) http.Handler {
		handler.next = next
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			handler.ServeHTTP(res, req)
		})
	}
}
