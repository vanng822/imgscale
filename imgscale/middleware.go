package imgscale

import (
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
				{"Prefix": "100x0", "Height": 100, "Ratio": 0.0, "Thumbnail": true},
				{"Prefix": "originalx1", "Height": 0, "Ratio": 1.0, "Thumbnail": false},
				{"Prefix": "original", "Height": 0, "Ratio": 0.0, "Thumbnail": false}
			],
			"Exts": ["jpg", "png"],
			"Comment": "Copyright"
		}
	
	
	The return handler could use as middleware handler
	
	Negroni middleware:
	
		n := negroni.New()
		n.UseHandler(imgscale.Configure("./config/formats.json"))
		http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8081), n)

	Martini middleware:
	
		app := martini.Classic()
		app.Use(imgscale.Configure("./config/formats.json").ServeHTTP)
		http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)

	http.Handle:
	
		http.Handle("/", imgscale.Configure("./config/formats.json"))
		http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8082), nil)

*/
func Configure(filename string) Handler {
	config := LoadConfig(filename)
	return configure(config)
}
