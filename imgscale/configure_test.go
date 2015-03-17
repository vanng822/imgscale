package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicExtNotSupported(t *testing.T) {
	config := &Config{}
	config.Exts = []string{"jpg", "bip"}
	assert.Panics(t, func() {
		configure(config)
	})
}

func TestPanicDuplicateFormat(t *testing.T) {
	config := &Config{}
	config.Exts = []string{"jpg", "png"}
	config.Formats = append(config.Formats, &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true})
	config.Formats = append(config.Formats, &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true})
	assert.Panics(t, func() {
		configure(config)
	})
}

func TestConfigureOK(t *testing.T) {
	config := &Config{}
	config.Path = "./"
	config.Prefix = "img"
	config.Exts = []string{"jpg", "png"}
	config.Formats = append(config.Formats, &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true})
	config.Formats = append(config.Formats, &Format{Prefix: "original", Height: 0, Ratio: 0.0, Thumbnail: false})
	handler := configure(config)
	expected := "^/img/(?P<format>100x100|original)/(?P<filename>.+\\.(?i)(jpg|png))$"
	assert.Equal(t, handler.regexp.String(), expected) 
	assert.Len(t, handler.formats, 2)
	assert.Equal(t, handler.supportedExts, supportedExts)
}

func TestConfigureSeparator(t *testing.T) {
	config := &Config{}
	config.Path = "./"
	config.Prefix = "img"
	config.Exts = []string{"jpg", "png"}
	config.Separator = "-"
	config.Formats = append(config.Formats, &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true})
	config.Formats = append(config.Formats, &Format{Prefix: "original", Height: 0, Ratio: 0.0, Thumbnail: false})
	handler := configure(config)
	expected := "^/img/(?P<format>100x100|original)-(?P<filename>.+\\.(?i)(jpg|png))$"
	assert.Equal(t, handler.regexp.String(), expected) 
	assert.Len(t, handler.formats, 2)
	assert.Equal(t, handler.supportedExts, supportedExts)
}

func TestConfigureNoExt(t *testing.T) {
	config := &Config{}
	config.Path = "./"
	config.Prefix = "img"
	config.Exts = []string{}
	config.Separator = "-"
	config.Formats = append(config.Formats, &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true})
	config.Formats = append(config.Formats, &Format{Prefix: "original", Height: 0, Ratio: 0.0, Thumbnail: false})
	handler := configure(config)
	expected := "^/img/(?P<format>100x100|original)-(?P<filename>.+)$"
	assert.Equal(t, handler.regexp.String(), expected) 
	assert.Len(t, handler.formats, 2)
	assert.Equal(t, handler.supportedExts, supportedExts)
}

func TestLoadConfig(t *testing.T) {
	config := LoadConfig("./test_config/formats.json")
	assert.Equal(t, config.conffile, "./test_config/formats.json")
	assert.Equal(t, config.Prefix, "img")
	assert.Equal(t, config.Path, "./test_data")
}
