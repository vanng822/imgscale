package middleware

import (
	"fmt"
	"github.com/go-martini/martini"
	"imgscale"
	"net/http"
	"strings"
	"strconv"
)

type Format struct {
	Prefix string
	Width  int
	Height int
	Keepratio bool
}

type Config struct {
	Path    string
	Prefix  string
	Formats []*Format
	Exts []string
	Comment string
}

var supportedExts = map[string]string{"jpg":"image/jpg", "png": "image/png"}

func Martini(config *Config, app *martini.ClassicMartini) {
	for _, ext := range config.Exts {
		if _, ok := supportedExts[ext]; !ok {
			panic(fmt.Sprintf("Extension '%s' not supported", ext))
		}
	}
	
	prefixes := make([]string, len(config.Formats))
	formats := make(map[string]*Format)
	for i, format := range config.Formats {
		prefixes[i] = format.Prefix
		formats[format.Prefix] = format
	}
	
	path := fmt.Sprintf("/%s/(?P<format>%s)/(?P<filename>.+).(?P<ext>%s)", config.Prefix, strings.Join(prefixes, "|"), strings.Join(config.Exts, "|"))
	fmt.Println(path)
	app.Get(path,
		func(c martini.Context, res http.ResponseWriter, req *http.Request, params martini.Params) {
			fmt.Println(params)
			format := formats[params["format"]]
			filename := fmt.Sprintf("%s/%s.%s", config.Path, params["filename"], params["ext"])
			info := imgscale.ImageInfo{filename, format.Width, format.Height, format.Keepratio}
			img, err := imgscale.GetImage(&info)
			if config.Comment != "" {
				img.CommentImage(config.Comment)
			}
			defer img.Destroy()
			if err == nil {
				imgData := img.GetImageBlob()
				res.Header().Set("Content-Type", supportedExts[params["ext"]])
				res.Header().Set("Content-Length", strconv.Itoa(len(imgData)))
				res.Write(imgData)
			} else {
				fmt.Println(err)
			}
		})
}
