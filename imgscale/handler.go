package imgscale

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

var supportedExts = map[string]string{"jpg": "image/jpeg", "png": "image/png"}

type Handler interface {
	// http.Handler
	ServeHTTP(res http.ResponseWriter, req *http.Request)
	// http.HandleFunc
	HandleFunc(res http.ResponseWriter, req *http.Request)
}

type handler struct {
	config        *Config
	formats       map[string]*Format
	regexp        *regexp.Regexp
	supportedExts map[string]string
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
	return &ImageInfo{fmt.Sprintf("%s/%s.%s", h.config.Path, filename, ext), f, ext, h.config.Comment}
}

func (h *handler) serve(res http.ResponseWriter, req *http.Request, info *ImageInfo) {
	img, err := GetImage(info)
	defer img.Destroy()
	if err == nil {
		imgData := img.GetImageBlob()
		res.Header().Set("Content-Type", h.getContentType(info.Ext))
		res.Header().Set("Content-Length", strconv.Itoa(len(imgData)))
		res.Write(imgData)
	} else {
		panic(err)
	}
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

func (h *handler) HandleFunc(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "HEAD" {
		return
	}
	matched, info := h.match(req.URL.RequestURI())
	if !matched {
		return
	}
	h.serve(res, req, info)	
}
