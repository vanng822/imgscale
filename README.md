## imgscale

Middleware/handler for scaling image in golang. Use for serving images in different formats. Can use as middleware which is compitable with http.HandleFunc

## Dependencies

	go get github.com/gographics/imagick/imagick

## Install 

	go get github.com/vanng822/imgscale/imgscale


## Example

	// Martini
	app := martini.Classic()
	app.Use(imgscale.Middleware("./config/formats.json"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)
	
	// Or http.HandleFunc
	http.HandleFunc("/", middleware.Middleware("./config/formats.json"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8080), nil)


## Try it out

### Checkout
	
	git clone https://github.com/vanng822/imgscale.git
	

### install dependencies

	go get github.com/gographics/imagick/imagick
	go get github.com/vanng822/imgscale/imgscale
	go get github.com/go-martini/martini
	
	
### run application
	go run main.go

### browse it
	
	http://127.0.0.1:8080/img/100x100/kth.jpg