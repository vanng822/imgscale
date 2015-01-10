package imgscale

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestGetImageWrongFile(t *testing.T) {
	filename, _ := filepath.Abs("./data/kth.jpg")
	f := &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true}
	info := &ImageInfo{filename, f, "jpg", ""}
	_, err := GetImage(info)
	assert.Error(t, err)
}

func TestGetImageScaleOK(t *testing.T) {
	filename, _ := filepath.Abs("../data/kth.jpg")
	f := &Format{Prefix: "133x100", Height: 100, Ratio: 0.0, Thumbnail: false}
	info := &ImageInfo{filename, f, "jpg", ""}
	img, err := GetImage(info)
	assert.Equal(t, 100, img.GetImageHeight())
	assert.Equal(t, 133, img.GetImageWidth())
	assert.Nil(t, err)
}

func TestGetImage100x100OK(t *testing.T) {
	filename, _ := filepath.Abs("../data/kth.jpg")
	f := &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true}
	info := &ImageInfo{filename, f, "jpg", ""}
	img, err := GetImage(info)
	assert.Equal(t, 100, img.GetImageHeight())
	assert.Equal(t, 100, img.GetImageWidth())
	assert.Nil(t, err)
}
