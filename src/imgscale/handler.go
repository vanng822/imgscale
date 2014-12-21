package imgscale

import (
	"net/http"
	//"time"
)

type ImageInfo struct {
	// full path to the image
	Filename  string
	Width     int
	Height    int
	KeepRatio bool
	// True if image should served original
	Original bool
}

type InfoParser interface {
	Parse(req *http.Request) *ImageInfo
}

type StandardInfoParser struct {
	SrcPath string
}

func (p StandardInfoParser) Parse(req *http.Request) *ImageInfo {
	// baseurl/img/{widthxheight}/image.{png,jpg}
	// example.com/img/24x23/test.jpg
	return &ImageInfo{"./data/kth.jpg", 100, 200, true, false}
}

var parser InfoParser

func Handler(res http.ResponseWriter, req *http.Request) {
	parser = StandardInfoParser{}
	info := parser.Parse(req)
	if info != nil {
		img := GetImage(info)
		res.Write(img.GetImageBlob())
	}
}
