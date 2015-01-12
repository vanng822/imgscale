package imgscale

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

var supportedExts = map[string]string{"jpg": "image/jpeg", "png": "image/png"}

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
	if h.validator != nil && h.validator.Validate(info.Filename) == false {
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
	return true, h.getImageInfo(matches[1], matches[2], matches[3])
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

	img, err := h.imageProvider.Fetch(info.Filename)
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
