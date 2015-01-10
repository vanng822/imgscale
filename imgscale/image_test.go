package imgscale

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestGetImageWrongFile(t *testing.T) {
	path, _ := filepath.Abs("./data/")
	provider := ImageProviderFile{path}
	
	filename := "kth.jpg"
	f := &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true}
	info := &ImageInfo{filename, f, "jpg", ""}
	_, err := provider.Fetch(info)
	assert.Error(t, err)
}

func TestGetImageScaleOK(t *testing.T) {
	path, _ := filepath.Abs("../data/")
	provider := ImageProviderFile{path}
	filename := "kth.jpg"
	f := &Format{Prefix: "133x100", Height: 100, Ratio: 0.0, Thumbnail: false}
	info := &ImageInfo{filename, f, "jpg", ""}
	img, err := provider.Fetch(info)
	assert.Nil(t, err)
	err = ProcessImage(img, info)
	assert.Equal(t, 100, img.GetImageHeight())
	assert.Equal(t, 133, img.GetImageWidth())
	assert.Nil(t, err)
}

func TestGetImage100x100OK(t *testing.T) {
	path, _ := filepath.Abs("../data/")
	provider := ImageProviderFile{path}
	filename := "kth.jpg"
	
	f := &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true}
	info := &ImageInfo{filename, f, "jpg", ""}
	img, err := provider.Fetch(info)
	assert.Nil(t, err)
	err = ProcessImage(img, info)
	assert.Equal(t, 100, img.GetImageHeight())
	assert.Equal(t, 100, img.GetImageWidth())
	assert.Nil(t, err)
}
