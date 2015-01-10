package imgscale

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Format struct {
	Prefix    string
	Height    uint
	Ratio     float64
	Thumbnail bool
}

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
		prefixes[i] = format.Prefix
		formats[format.Prefix] = format
	}

	path := fmt.Sprintf("/%s/(?P<format>%s)/(?P<filename>.+)\\.(?P<ext>%s)", config.Prefix, strings.Join(prefixes, "|"), strings.Join(config.Exts, "|"))

	return &handler{formats: formats, config: config, regexp: regexp.MustCompile(path), supportedExts: supportedExts}
}

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
