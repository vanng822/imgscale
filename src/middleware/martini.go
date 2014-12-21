package middleware

import (
	"fmt"
	"github.com/go-martini/martini"
	"imgscale"
	"net/http"
	"strings"
)

type Format struct {
	Prefix string
	Width  int
	Height int
}

type Config struct {
	Path    string
	Prefix  string
	Formats []*Format
	Exts []string
}

func Martini(config *Config, app *martini.ClassicMartini) {
	prefixes := make([]string, len(config.Formats))
	formats := make(map[string]*Format)
	for i, format := range config.Formats {
		prefixes[i] = format.Prefix
		formats[format.Prefix] = format
	}
	path := fmt.Sprintf("/%s/(?P<format>%s)/(?P<filename>.+.(%s))", config.Prefix, strings.Join(prefixes, "|"), strings.Join(config.Exts, "|"))
	fmt.Println(path)
	app.Get(path,
		func(c martini.Context, res http.ResponseWriter, req *http.Request, params martini.Params) {
			fmt.Println(params)
			format := formats[params["format"]]
			filename := fmt.Sprintf("%s/%s", config.Path, params["filename"])
			info := imgscale.ImageInfo{filename, format.Width, format.Height, true}
			img, err := imgscale.GetImage(&info)
			defer img.Destroy()
			if err == nil {
				res.Write(img.GetImageBlob())
			} else {
				fmt.Println(err)
			}
		})
}
