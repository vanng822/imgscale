package imgscale

import (
	"github.com/gographics/imagick/imagick"
	"net/http"
)

type ImageProvider interface {
	Fetch(filename string) (*imagick.MagickWand, error)
}

type Validator interface {
	// Name of the image, ie everything after "<prefix>/<format>/"
	Validate(filename string) bool
}

// http.Handler
type Handler interface {
	// http.HandleFunc
	ServeHTTP(res http.ResponseWriter, req *http.Request)
	// For setting own image provider, such as remote storage
	SetImageProvider(provider ImageProvider)
	// For setting validator of filename/name of the image
	SetValidator(validator Validator)
}

