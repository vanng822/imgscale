## imgscale

Middleware/handler for scaling image in golang. Use for serving images in different formats. Can use as middleware which is compitable with http.HandleFunc

## Dependencies

	go get github.com/gographics/imagick/imagick

## Example

	// Martini
	app := martini.Classic()
	app.Use(imgscale.Middleware("./config/formats.json"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)
	
	// http.HandleFunc
	http.HandleFunc("/", middleware.Middleware("./config/formats.json"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8080), nil)
