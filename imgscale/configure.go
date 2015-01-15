package imgscale

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
	Prefix: use in the url to identify the format
	
	Height: the target height
	
	Ratio: zero for keeping ratio
	
	Thumbnail: true to use thumbnail feature in imagemagick, it is quick and optimized but you loose meta data
	
*/
type Format struct {
	Prefix    string
	Height    uint
	Ratio     float64
	Thumbnail bool
}

/*
	Path: path of the folder which contains images
	
	Prefix: use as image path indicator in url
	
	Formats: list of Format
	
	Exts: allow extensions, only jpg and png available
	
	Comment: will store in meta data if specified
	
*/
type Config struct {
	Path    string
	Prefix  string
	Formats []*Format
	Exts    []string
	Comment string
}

func configure(config *Config) *handler {
	for _, ext := range config.Exts {
		if _, ok := supportedExts[ext]; !ok {
			panic(fmt.Sprintf("Extension '%s' not supported", ext))
		}
	}

	prefixes := make([]string, len(config.Formats))
	formats := make(map[string]*Format)
	for i, format := range config.Formats {
		if _, exists := formats[format.Prefix]; exists {
			panic(fmt.Sprintf("You can not have same prefix '%s' for 2 different formats", format.Prefix))
		}
		prefixes[i] = format.Prefix
		formats[format.Prefix] = format
	}

	path := fmt.Sprintf("/%s/(?P<format>%s)/(?P<filename>.+)\\.(?P<ext>%s)", config.Prefix, strings.Join(prefixes, "|"), strings.Join(config.Exts, "|"))
	h := handler{formats: formats, config: config, regexp: regexp.MustCompile(path), supportedExts: supportedExts}
	h.SetValidator(defaultValidator{})
	return &h
}

// LoadConfig parse configuration file
// It will panic if any error
func LoadConfig(filename string) *Config {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
