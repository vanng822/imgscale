package imgscale

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var supportedExts = map[string]string{"jpg": "image/jpeg", "png": "image/png"}

type ImageProvider interface {
	Fetch(info *ImageInfo) (*imagick.MagickWand, error)
}

type Validator interface {
	// Name of the image, ie everything after "<prefix>/<format>/"
	Validate(filename string) bool
}

type defaultValidator struct {}

func (v defaultValidator) Validate(filename string) bool {
	return strings.Index(filename, "..") == -1
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

type handler struct {
	config        *Config
	formats       map[string]*Format
	regexp        *regexp.Regexp
	supportedExts map[string]string
	imageProvider ImageProvider
	validator     Validator
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "HEAD" {
		return
	}
	matched, info := h.match(req.URL.RequestURI())
	if !matched {
		return
	}
	h.serve(res, req, info)
}

func (h *handler) SetImageProvider(imageProvider ImageProvider) {
	h.imageProvider = imageProvider
}

func (h *handler) SetValidator(validator Validator) {
	h.validator = validator
}

func (h *handler) match(url string) (bool, *ImageInfo) {
	matches := h.regexp.FindStringSubmatch(url)

	if len(matches) == 0 {
		return false, nil
	}
	info := h.getImageInfo(matches[1], matches[2], matches[3])
	if h.validator != nil && h.validator.Validate(info.Filename) == false {
		return false, nil
	}
	return true, info
}

func (h *handler) getContentType(ext string) string {
	return h.supportedExts[ext]
}

func (h *handler) getFormat(format string) *Format {
	return h.formats[format]
}

func (h *handler) getImageInfo(format, filename, ext string) *ImageInfo {
	f := h.getFormat(format)
	return &ImageInfo{fmt.Sprintf("%s.%s", filename, ext), f, ext, h.config.Comment}
}

func (h *handler) serve(res http.ResponseWriter, req *http.Request, info *ImageInfo) {
	if h.imageProvider == nil {
		h.imageProvider = NewImageProviderFile(h.config.Path)
	}

	img, err := h.imageProvider.Fetch(info)
	if img != nil {
		defer img.Destroy()
	}
	if err != nil {
		return
	}
	err = ProcessImage(img, info)
	if err == nil {
		imgData := img.GetImageBlob()
		res.Header().Set("Content-Type", h.getContentType(info.Ext))
		res.Header().Set("Content-Length", strconv.Itoa(len(imgData)))
		res.Write(imgData)
	}
}
