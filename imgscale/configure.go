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

	Watermark: Indicate if watermark should apply on this format
	
	Quality: Compression quality
	
	Strip: strips all profiles and original comments

*/
type Format struct {
	Prefix    string
	Height    uint
	Ratio     float64
	Thumbnail bool
	Watermark bool
	Quality   uint
	Strip     bool
}

/*
	Path: path of the folder which contains images

	Prefix: use as image path indicator in url

	Separator: separator between format prefix and filename

	Formats: list of Format

	Exts: allow extensions, only jpg, jpeg and png available. Be aware that jpg and jpeg handled as different extension.

	Comment: will store in meta data if specified

	AutoRotate: Autorotate image according to orientation stored in the image meta data

	Watermark: this watermark is used when watermark on a format is enabled

*/
type Config struct {
	Path       string
	Prefix     string
	Separator  string
	Formats    []*Format
	Exts       []string
	Comment    string
	AutoRotate bool
	Watermark  *Watermark
	conffile   string
}

func checkExtension(config *Config) {
	for _, ext := range config.Exts {
		if _, ok := supportedExts[ext]; !ok {
			panic(fmt.Sprintf("Extension '%s' not supported", ext))
		}
	}
}

func getFormats(config *Config) map[string]*Format {
	formats := make(map[string]*Format)
	for _, format := range config.Formats {
		if _, exists := formats[format.Prefix]; exists {
			panic(fmt.Sprintf("You can not have same prefix '%s' for 2 different formats", format.Prefix))
		}
		formats[format.Prefix] = format
	}
	return formats
}

func compilePath(config *Config) *regexp.Regexp {
	var separator string
	if config.Separator != "" {
		separator = config.Separator
	} else {
		separator = "/"
	}
	prefixes := make([]string, 0)
	for _, format := range config.Formats {
		prefixes = append(prefixes, format.Prefix)
	}
	path := fmt.Sprintf("^/%s/(?P<format>%s)%s(?P<filename>.+)\\.(?i)(?P<ext>%s)$",
		config.Prefix,
		strings.Join(prefixes, "|"),
		separator,
		strings.Join(config.Exts, "|"))

	return regexp.MustCompile(path)
}

// setting up stuffs based on config
func setupHandlerConfig(h *handler, config *Config) {
	checkExtension(config)
	formats := getFormats(config)
	regexp := compilePath(config)
	if config.Watermark != nil && config.Watermark.Filename != "" {
		config.Watermark.load()
	}
	h.formats = formats
	h.regexp = regexp
	h.config = config
}

func configure(config *Config) *handler {
	h := new(handler)
	h.supportedExts = supportedExts
	h.SetValidator(defaultValidator{})
	setupHandlerConfig(h, config)
	return h
}

/*
	LoadConfig parse configuration file. It will panic if any error
*/
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
	conf.conffile = filename
	return &conf
}
