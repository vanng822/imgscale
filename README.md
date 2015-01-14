## imgscale

Middleware/handler for scaling image in golang. Use for serving images in different formats. Can use as middleware which is compitable with http.Handler

Warning: image processing is very resource consuming. If you use this in production you should put a cache server in front, such as Varnish.

## Dependencies

	go get github.com/gographics/imagick/imagick

## Install 

	go get github.com/vanng822/imgscale/imgscale


## Example

	// Negroni
	n := negroni.New()
	n.UseHandler(imgscale.Configure("./config/formats.json"))
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8081), n)

	// Martini
	app := martini.Classic()
	app.Use(imgscale.Configure("./config/formats.json").ServeHTTP)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)

	// http.Handler
	http.Handle("/", imgscale.Configure("./config/formats.json"))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8082), nil)

## GoDoc

[![GoDoc](https://godoc.org/github.com/vanng822/imgscale/imgscale?status.svg)](https://godoc.org/github.com/vanng822/imgscale/imgscale)


## Try it out

### checkout
	
	git clone https://github.com/vanng822/imgscale.git
	

### install dependencies

	go get github.com/gographics/imagick/imagick
	go get github.com/vanng822/imgscale/imgscale
	go get github.com/go-martini/martini
	go get github.com/codegangsta/negroni
	
	
### run application

	go run main.go

### browse it
	
	http://127.0.0.1:8080/img/100x100/kth.jpg
	http://127.0.0.1:8081/img/100x100/http://127.0.0.1:8080/img/original/kth.jpg
	http://127.0.0.1:8082/img/100x100/kth.jpg
	