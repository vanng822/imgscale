package imgscale

import (
	"github.com/vanng822/imgscale/imagick"
	"net/http"
)

/*
	ImageProvider implements image fetching and provide the image source
	for the handler to serve the request. Local filesystem imageProviderFile
	is bundled but there are more providers. Take look at provider folder
	Fetch has to return MagickWand in favour of imageProviderFile. This case
	we read image data directly to MagickWand
*/
type ImageProvider interface {
	Fetch(filename string) (*imagick.MagickWand, error)
}

/*
	Validator implements the validation of the filename
	The filename can identified as
	baseurl/<image prefix>/<format prefix><separator><filename>
	An example is
	http://127.0.0.1:8080/img/original/kth.jpg
	and filename is kth.jpg
*/
type Validator interface {
	// Name of the image, ie everything after "<prefix>/<format><separator>"
	Validate(filename string) bool
}

/* 
	Handler implements http.Handler so it can use for many frameworks available
	Handler.ServeHTTP can use similar to http.HandleFunc in case frameworks support
	only this.
	
	Beside that Handler has some more methods, Handler.SetValidator for setting own validation of the filename
	Handler.SetImageProvider is suitable when you have customized image provider, default is
	imageProviderFile. And the last method Handler.Cleanup should always call at the end (or defer) to cleanup C pointers.
	Reload will reload new configurations without stopping service
*/
type Handler interface {
	// http.HandleFunc
	ServeHTTP(res http.ResponseWriter, req *http.Request)
	// For setting own image provider, such as remote storage
	SetImageProvider(provider ImageProvider)
	// For setting validator of filename/name of the image
	SetValidator(validator Validator)
	// Free C pointers and terminate MagickWand environment
	Cleanup()
	// Reload configuration
	Reload()
	// func(next http.Handler) http.Handler
	Middleware() func(next http.Handler) http.Handler
	// HandlerFunc With Next HandlerFunc
	HandlerFuncWithNext() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

